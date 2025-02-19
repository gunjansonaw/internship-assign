package routes

import (
	"holiday-calendar/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterHolidayRoutes(router *gin.Engine) {
	holidayRoutes := router.Group("/api/holidays")
	{
		holidayRoutes.POST("/", controllers.AddHoliday)
		holidayRoutes.GET("/", controllers.GetHolidays)
		holidayRoutes.DELETE("/:id", controllers.DeleteHoliday)
	}
}
