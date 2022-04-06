package major

import (
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"pmc_server/init/aura"
)

type Entity struct {
	CollegeID        int
	Name             string
	DegreeHour       int32
	MinMajorHour     int32
	EmphasisRequired bool
}

type Insertion struct {
	Major Entity
}

func (m Insertion) InsertMajor() (string, error) {
	session := aura.Driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, err := session.WriteTransaction(m.insertMajorFn)
	if err != nil {
		return "", err
	}
	return result.(string), nil
}

func (m *Insertion) insertMajorFn(tx neo4j.Transaction) (interface{}, error) {
	records, err := tx.Run("CREATE (m:Major { name: $name, degree_hour: $degree_hour,"+
		" min_major_hour: $min_major_hour, emphasis_required: $emphasis_required, college_id: $college_id}) "+
		"RETURN m.college_id, m.name, m.degree_hour, m.min_major_hour, m.emphasis_required",
		map[string]interface{}{
			"college_id":        m.Major.CollegeID,
			"name":              m.Major.Name,
			"degree_hour":       m.Major.DegreeHour,
			"min_major_hour":    m.Major.MinMajorHour,
			"emphasis_required": m.Major.EmphasisRequired,
		})

	if err != nil {
		return nil, err
	}

	record, err := records.Single()
	if err != nil {
		return nil, err
	}

	return record.Values[1], nil
}

type Emphasis struct {
	Name          string
	TotalCredit   int32
	MainMajorName string
	CollegeID     int32
}

type EmphasisInsertion struct {
	Emphasis Emphasis
}

func (m EmphasisInsertion) InsertEmphasis() (string, error) {
	session := aura.Driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, err := session.WriteTransaction(m.insertEmphasisFn)
	if err != nil {
		return "", err
	}
	return result.(string), nil
}

func (m *EmphasisInsertion) insertEmphasisFn(tx neo4j.Transaction) (interface{}, error) {
	records, err := tx.Run("MATCH (m:Emphasis) WHERE m.m = $major_name "+
		"CREATE (e:Emphasis)-[:SUB_OF]->(m) SET e.name = $name, e.total_credit = $total_credit, "+
		"e.college_id = $college_id RETURN e.name",
		map[string]interface{}{
			"major_name":   m.Emphasis.MainMajorName,
			"name":         m.Emphasis.Name,
			"total_credit": m.Emphasis.TotalCredit,
			"college_id":   m.Emphasis.CollegeID,
		})

	if err != nil {
		return nil, err
	}

	record, err := records.Single()
	if err != nil {
		return nil, err
	}

	return record.Values[0], nil
}

type DegreeType struct {
	Name      string
	Major     string
	CollegeID int32
}

type DegreeInsertion struct {
	Type DegreeType
}

func (m DegreeInsertion) InsertDegreeType() (string, error) {
	session := aura.Driver.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	result, err := session.WriteTransaction(m.insertDegreeTypeFn)
	if err != nil {
		return "", err
	}
	return result.(string), nil
}

func (m *DegreeInsertion) insertDegreeTypeFn(tx neo4j.Transaction) (interface{}, error) {
	records, err := tx.Run("MATCH (m:Major) WHERE m.name = $major_name "+
		"CREATE (d:DegreeType)<-[HAS_DEGREE]-(m) SET d.name = $name, d.college_id = $college_id RETURN d.name",
		map[string]interface{}{
			"major_name": m.Type.Major,
			"name":       m.Type.Name,
			"college_id": m.Type.CollegeID,
		})
	if err != nil {
		return nil, err
	}

	record, err := records.Single()
	if err != nil {
		return nil, err
	}

	return record.Values[0], nil
}

type Read struct {
	CollegeID int32
}

func (r Read) FindAll() ([]Entity, error) {
	session := aura.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	result, err := session.ReadTransaction(r.findAllFn)
	if err != nil {
		return nil, err
	}

	return result.([]Entity), nil
}

func (r *Read) findAllFn(tx neo4j.Transaction) (interface{}, error) {
	res, err := tx.Run("MATCH (m:Major {college_id : $college_id}) "+
		"RETURN m.degree_hour, m.emphasis_required, m.min_major_hour, m.name",
		map[string]interface{}{
			"college_id": r.CollegeID,
		})

	if err != nil {
		return nil, err
	}

	majorList := make([]Entity, 0)
	for res.Next() {
		var major Entity
		if name, ok := res.Record().Values[3].(string); ok {
			major.Name = name
		}
		if degreeHour, ok := res.Record().Values[0].(int64); ok {
			major.DegreeHour = int32(degreeHour)
		}
		if minMajorHour, ok := res.Record().Values[2].(int64); ok {
			major.MinMajorHour = int32(minMajorHour)
		}
		if emphasisRequired, ok := res.Record().Values[1].(bool); ok {
			major.EmphasisRequired = emphasisRequired
		}

		majorList = append(majorList, major)
	}

	return majorList, nil
}

type ReadEmphasis struct {
	CollegeID int32
	MajorName string
}

func (r ReadEmphasis) FindAllEmphasisesOfAMajor() ([]Entity, error) {
	session := aura.Driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
	defer session.Close()

	result, err := session.ReadTransaction(r.findAllEmphasisesFn)
	if err != nil {
		return nil, err
	}

	return result.([]Entity), nil
}

func (r *ReadEmphasis) findAllEmphasisesFn(tx neo4j.Transaction) (interface{}, error) {
	res, err := tx.Run("MATCH (m:Major {name: $major_name, college_id: $college_id})<-[:SUB_OF]-(emphasis) "+
		"RETURN emphasis.name",
		map[string]interface{}{
			"major_name": r.MajorName,
			"college_id": r.CollegeID,
		})

	if err != nil {
		return nil, err
	}

	emphasisList := make([]Emphasis, 0)
	for res.Next() {
		var emphasis Emphasis
		if name, ok := res.Record().Values[0].(string); ok {
			emphasis.Name = name
		}
		emphasisList = append(emphasisList, emphasis)
	}

	return emphasisList, nil
}
