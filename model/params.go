package model

type RegisterParams struct {
	Email      string `json:"email"`      // User Email
	FirstName  string `json:"firstName"`  // User first name
	LastName   string `json:"lastName"`   // User Last name
	College    string `json:"college"`    // User college, this is restricted to UofU now
	Password   string `json:"password"`   // User provided password, will be encrypted
	RePassword string `json:"rePassword"` // User reentered password, to make sure they match
}

type LoginParams struct {
	Email    string `json:"email" binding:"required"`    // User Email
	Password string `json:"password" binding:"required"` // User password
}

type CourseParams struct {
	ID string `uri:"id" binding:"required"` // Course ID
}

type CourseFilterParams struct {
	// TODO: fix this with actual professor entities
	TaughtProfessor       []string `json:"taughtProfessor"`       // Filter courses with professor names
	Keyword               string   `json:"keyword"`               // Keyword user inputs, this links to the name/catalog name/subject/tag of the course
	MinCredit             float32  `json:"minCredit"`             // Filter courses with given credit
	MaxCredit             float32  `json:"maxCredit"`             //Filter cores with the given max credit
	OfferedOnline         bool     `json:"isOnline"`              // Filter courses that's online
	OfferedOffline        bool     `json:"isOffline"`             // Filter courses that's in person
	OfferedIVC            bool     `json:"offered_ivc"`           // Filter classes that's ivc
	OfferedHybrid         bool     `json:"offered_hybrid"`        // Filter classes that's hybrid
	IsHonor               bool     `json:"isHonor"`               // Filter courses that's honor courses
	Weekday               []int    `json:"weekday"`               // Filter courses that's in the specific weekdays
	StartTime             float32  `json:"startTime"`             // Filter courses that starts no earlier than the start time
	EndTime               float32  `json:"endTime"`               // Filter courses that ends no later than the start time
	MinRating             int8     `json:"minRating"`             // Filter courses that has no lower rating than the give min rating
	RankByRatingHighToLow bool     `json:"rankByRatingHighToLow"` // Rank the courses by the given rating low to high
	RankByRatingLowToHigh bool     `json:"rankByRatingLowToHigh"` // Rank the courses by the given rating low to high
	PageNumber            int      `json:"pageNumber"`            // The current page of the search result, default 0
	PageSize              int      `json:"pageSize"`              // The current page size of the search result, default 10
}

type ClassParams struct {
	ID string `uri:"id" binding:"required"` // Class ID
}

type ReviewParams struct {
	CourseID    int64   `json:"courseID" binding:"required"`    // Course ID
	UserID      int64   `json:"userID" binding:"required"`      // User ID
	Pros        string  `json:"pros" binding:"required"`        // The pros of the course given by the user
	Cons        string  `json:"cons" binding:"required"`        // The cons of the course given by the user
	Comment     string  `json:"comments" binding:"required"`    // The detailed comment on the review given by the user
	Rating      float32 `json:"rating" binding:"required"`      // The rating of the course given by the user
	IsAnonymous bool    `json:"isAnonymous" binding:"required"` // Is the user posting this review anonymously
	Recommended bool    `json:"recommended" binding:"required"` // Is the course recommended by the user
}

type PostScheduleParams struct {
	ClassID    int64 `json:"classID"`    // Class ID
	SemesterID int64 `json:"semesterID"` // Semester ID
	UserID     int64 `json:"userID"`     // Student ID
}

type GetScheduleParams struct {
	UserID int64 `form:"userID"` // Student ID
}

type DeleteScheduleParams struct {
	UserID     int64 `json:"userID"`     // Student ID
	ClassID    int64 `json:"classID"`    // Class ID
	SemesterID int64 `json:"semesterID"` // Semester ID
}

type CreateTagParam struct {
	CourseID int64  `json:"courseID"` // The course ID
	Content  string `json:"content"`  // The content of the tag
}

type VoteTagParam struct {
	TagID  int32 `json:"tagID"`  // The Tag ID
	UserID int64 `json:"userID"` // The user ID
	Upvote bool  `json:"upvote"` // Check if the vote is an upvote or down vote, 1 for up, 0 for down
}
