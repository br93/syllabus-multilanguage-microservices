package controllers

import (
	"net/http"
	"strings"

	"github.com/br93/syllabus/syllabus-settings-go/pkg/mappers"
	"github.com/br93/syllabus/syllabus-settings-go/pkg/models"
	"github.com/br93/syllabus/syllabus-settings-go/pkg/services"
	"github.com/gin-gonic/gin"
)

var NewCurso models.CursoEntity

func CreateCurso(ctx *gin.Context) {
	body := models.CursoRequestModel{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	curso := mappers.ToCursoEntity(&body)

	if err := services.CreateCurso(curso); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response := mappers.ToCursoResponse(curso)

	ctx.JSON(http.StatusCreated, response)
}

func GetCursoById(ctx *gin.Context) {
	cursoId := ctx.Param("curso_id")

	curso, err := services.GetCursoById(cursoId)

	if err != nil && strings.Contains(err.Error(), "not found") {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	} else if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response := mappers.ToCursoResponse(curso)

	ctx.JSON(http.StatusOK, response)

}

func GetCursos(ctx *gin.Context) {

	cursos, err := services.GetCursos()

	if err != nil && strings.Contains(err.Error(), "not found") {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	} else if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response := mappers.ToCursoResponseArray(cursos)

	ctx.JSON(http.StatusOK, response)
}

func UpdateCurso(ctx *gin.Context) {
	cursoId := ctx.Param("curso_id")
	body := models.CursoRequestModel{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	curso := mappers.ToCursoEntity(&body)
	update, err := services.UpdateCurso(cursoId, curso)

	if err != nil && strings.Contains(err.Error(), "not found") {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	} else if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response := mappers.ToCursoResponse(update)

	ctx.JSON(http.StatusOK, response)
}

func DeleteCurso(ctx *gin.Context) {
	cursoId := ctx.Param("curso_id")

	err := services.DeleteCurso(cursoId)
	if err != nil && strings.Contains(err.Error(), "not found") {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	} else if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
