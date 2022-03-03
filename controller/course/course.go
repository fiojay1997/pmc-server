package controller

import (
	"net/http"

	. "pmc_server/consts"
	"pmc_server/logic"
	"pmc_server/model"
	"pmc_server/utils"

	"github.com/gin-gonic/gin"
)

// GetCourseListHandler Get the entire course list
// @Summary Use this API to get the list of the classes
// @Description This API is used to get the course list, you should do pagination
// @Tags Course
// @Accept application/json
// @Produce application/json
// @Param pn query int false "Page number of the paginated course list, default is 1" mininum(1)
// @Param psize query int false "Page size(number of course to fetch) of the paginated course list, default is 10" mininum(1) maxinum(30)
// @Success 200 {int} total number of the courses
// @Success 200 {array} dto.Course
// @Router /course/list [get]
func GetCourseListHandler(c *gin.Context) {
	pnInt, pSizeInt, err := utils.HandlePagination(c, "10")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ERROR: err,
		})
		return
	}

	courseList, total, err := logic.GetCourseList(pnInt, pSizeInt)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			ERROR: err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		DATA:  courseList,
		TOTAL: total,
	})
}

// GetCourseByIDHandler Get course and its classes by the given ID
// @Summary Use this API to get the class by the given ID
// @Description This API is used to get the course info along with the classes
// @Tags Course
// @Accept application/json
// @Produce application/json
// @Param id query int true "course id"
// @Success 200 {object} dto.Class
// @Router /course/:id [get]
func GetCourseByIDHandler(c *gin.Context) {
	var courseParam model.CourseParams
	if err := c.ShouldBindUri(&courseParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			MESSAGE: NO_ID_ERR,
		})
		return
	}

	courseInfo, err := logic.GetCourseInfo(courseParam.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			MESSAGE: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		DATA: &courseInfo,
	})
}

// GetClassesOfCourseHandler Get the class list of the given course
// @Summary Use this API to get the list of the classes
// @Description This API is used to get the class list, you should do pagination
// @Tags Course
// @Accept application/json
// @Produce application/json
// @Param id query int true "course id"
// @Success 200 {int} total number of the courses
// @Success 200 {array} dto.Class
// @Router /course/:id/class [get]
func GetClassesOfCourseHandler(c *gin.Context) {
	var courseParam model.CourseParams
	if err := c.ShouldBindUri(&courseParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			MESSAGE: NO_ID_ERR,
		})
		return
	}

	classList, total, err := logic.GetClassListByCourseID(courseParam.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			MESSAGE: NO_INFO_ERR,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		MESSAGE: SUCCESS,
		TOTAL:   total,
		DATA:    classList,
	})
}

// GetCoursesBySearchHandler Get the entire course list
// @Summary Use this API to get the list of the classes
// @Description This API is used to get the course list, you should do pagination
// @Tags Course
// @Accept application/json
// @Produce application/json
// @Success 200 {int} total number of the courses
// @Success 200 {array} dto.Course
// @Router /course/list [post]
func GetCoursesBySearchHandler(c *gin.Context) {
	var param model.CourseFilterParams
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			ERROR: INSUFFICIENT_PARAM_ERR,
		})
		return
	}

	data, err := logic.GetCoursesBySearch(param)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			ERROR: NO_ID_ERR,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		DATA: data,
	})
}
