package controllers

import (
	"net/http"
	"strings"

	"github.com/br93/syllabus/syllabus-settings-go/pkg/mappers"
	"github.com/br93/syllabus/syllabus-settings-go/pkg/models"
	"github.com/br93/syllabus/syllabus-settings-go/pkg/services"
	"github.com/br93/syllabus/syllabus-settings-go/pkg/utils"
	"github.com/gin-gonic/gin"
)

var NewCourseProgram models.CourseProgram

func CreateCourseProgram(ctx *gin.Context) {
	body := models.CourseProgramRequestModel{}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ErrorHandling(ctx, err)
		return
	}

	courseprogram := mappers.ToCourseProgram(&body)

	if err := services.CreateCourseProgram(courseprogram); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response := mappers.ToCourseProgramResponse(courseprogram)

	ctx.JSON(http.StatusCreated, response)
}

func GetCourseProgramById(ctx *gin.Context) {
	courseprogramId := ctx.Param("course_program_id")

	courseprogram, err := services.GetCourseProgramById(courseprogramId)

	if err != nil && strings.Contains(err.Error(), "not found") {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	} else if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response := mappers.ToCourseProgramResponse(courseprogram)

	ctx.JSON(http.StatusOK, response)

}

func GetCoursePrograms(ctx *gin.Context) {

	courseprograms, err := services.GetCoursePrograms("Course", "CourseType")

	if err != nil && strings.Contains(err.Error(), "not found") {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	} else if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response := mappers.ToCourseProgramResponseArray(courseprograms)

	ctx.JSON(http.StatusOK, response)
}

func GetCourseProgramsByCourseCodeIn(ctx *gin.Context) {

	codes, _ := ctx.GetQueryArray("a")

	courseprograms, err := services.GetCourseProgramsByCourseCodeIn(codes, "Course", "CourseType")

	if err != nil && strings.Contains(err.Error(), "not found") {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	} else if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response := mappers.ToCourseProgramResponseArray(courseprograms)

	ctx.JSON(http.StatusOK, response)
}

func GetCourseProgramsByProgramAndCourseType(ctx *gin.Context) {

	programCode := ctx.Param("program_code")
	typeName := ctx.Param("type_name")

	courseprograms, err := services.GetCourseProgramsByProgramAndCourseType(programCode, typeName, "Course", "CourseType")

	if err != nil && strings.Contains(err.Error(), "not found") {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	} else if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response := mappers.ToCourseProgramResponseArray(courseprograms)

	ctx.JSON(http.StatusOK, response)
}

func UpdateCourseProgram(ctx *gin.Context) {
	courseprogramId := ctx.Param("course_program_id")
	body := models.CourseProgramRequestModel{}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ErrorHandling(ctx, err)
		return
	}

	courseprogram := mappers.ToCourseProgram(&body)
	update, err := services.UpdateCourseProgram(courseprogramId, courseprogram)

	if err != nil && strings.Contains(err.Error(), "not found") {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	} else if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response := mappers.ToCourseProgramResponse(update)

	ctx.JSON(http.StatusOK, response)
}

func DeleteCourseProgram(ctx *gin.Context) {
	courseprogramId := ctx.Param("course_program_id")

	err := services.DeleteCourseProgram(courseprogramId)
	if err != nil && strings.Contains(err.Error(), "not found") {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	} else if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
