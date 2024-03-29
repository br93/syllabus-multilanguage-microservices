package middlewares

import (
	"net/http"

	"github.com/br93/syllabus/syllabus-settings-go/pkg/cache"
	"github.com/br93/syllabus/syllabus-settings-go/pkg/unmarshal"
	"github.com/gin-gonic/gin"
)

func CacheClassSchedules(ctx *gin.Context) {
	classSchedules := cache.Get("all-class-schedules")

	if classSchedules != "nil" {
		response := unmarshal.UnmarshalClassSchedules(classSchedules)
		ctx.JSON(http.StatusOK, response)
		ctx.Abort()
	}

	ctx.Next()
}

func CacheClassSchedule(ctx *gin.Context) {
	classScheduleId := ctx.Param("class_schedule_id")

	classSchedule := cache.Get("class-schedule" + classScheduleId)

	if classSchedule != "nil" {
		response := unmarshal.UnmarshalClassSchedule(classSchedule)
		ctx.JSON(http.StatusOK, response)
		ctx.Abort()
	}

	ctx.Next()
}

func CacheClassSchedulesByClass(ctx *gin.Context) {
	classId := ctx.Param("class_id")

	classSchedules := cache.Get("class-schedules" + classId)

	if classSchedules != "nil" {
		response := unmarshal.UnmarshalClassSchedules(classSchedules)
		ctx.JSON(http.StatusOK, response)
		ctx.Abort()
	}

	ctx.Next()
}
