package main

import (
	"log"

	"github.com/cyneptic/letsgo_smspanel_mockapi/controller"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	
	controller.AddMessageRoutes(e)

	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}