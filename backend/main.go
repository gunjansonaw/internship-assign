package main

import (
	"fmt"
	"holiday-calendar/config"
	"holiday-calendar/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.ConnectDatabase()

	routes.RegisterHolidayRoutes(router)

	fmt.Println("Server running on port 8080")
	router.Run(":8080")
}
