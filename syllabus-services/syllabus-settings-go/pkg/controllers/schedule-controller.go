package controllers

import (
	"net/http"
	"strings"

	"github.com/br93/syllabus/syllabus-settings-go/pkg/cache"
	"github.com/br93/syllabus/syllabus-settings-go/pkg/mappers"
	"github.com/br93/syllabus/syllabus-settings-go/pkg/models"
	"github.com/br93/syllabus/syllabus-settings-go/pkg/services"
	"github.com/br93/syllabus/syllabus-settings-go/pkg/utils"
	"github.com/gin-gonic/gin"
)

var NewSchedule models.Schedule

func CreateSchedule(ctx *gin.Context) {
	body := models.ScheduleRequestModel{}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ErrorHandling(ctx, err)
		return
	}

	schedule := mappers.ToSchedule(&body)

	if err := services.CreateSchedule(schedule); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response := mappers.ToScheduleResponse(schedule)

	cache.Flush()
	ctx.JSON(http.StatusCreated, response)
}

func GetScheduleByIdOrCode(ctx *gin.Context) {
	scheduleId := ctx.Param("schedule_id")

	schedule, err := services.GetScheduleByIdOrCode(scheduleId)

	if err != nil && strings.Contains(err.Error(), "not found") {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	} else if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response := mappers.ToScheduleResponse(schedule)
	cache.Set("schedule", scheduleId, response)

	ctx.JSON(http.StatusOK, response)

}

func GetSchedules(ctx *gin.Context) {

	schedules, err := services.GetSchedules()

	if err != nil && strings.Contains(err.Error(), "not found") {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	} else if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response := mappers.ToScheduleResponseArray(schedules)
	cache.SetAll("all-schedules", response)

	ctx.JSON(http.StatusOK, response)
}

func UpdateSchedule(ctx *gin.Context) {
	scheduleId := ctx.Param("schedule_id")
	body := models.ScheduleRequestModel{}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ErrorHandling(ctx, err)
		return
	}

	schedule := mappers.ToSchedule(&body)
	update, err := services.UpdateSchedule(scheduleId, schedule)

	if err != nil && strings.Contains(err.Error(), "not found") {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	} else if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response := mappers.ToScheduleResponse(update)

	cache.Flush()
	ctx.JSON(http.StatusOK, response)
}

func DeleteSchedule(ctx *gin.Context) {
	scheduleId := ctx.Param("schedule_id")

	err := services.DeleteSchedule(scheduleId)
	if err != nil && strings.Contains(err.Error(), "not found") {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	} else if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	cache.Flush()
	ctx.JSON(http.StatusNoContent, nil)
}
