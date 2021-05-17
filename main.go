package main

import (
	"fmt"

	"acgfate/config"
	"acgfate/router"
)

// @title ACG.Fate API
// @version 0.0.1
// @description The RESTFul API of Server
// @host 127.0.0.1:3000
// @BasePath /api/v1
func main() {
	// Read config file
	config.ReadConfig()

	// Initialize the router
	r := router.InitRouter()

	// Run server
	err := r.Run(":3000")
	if err != nil {
		panic(fmt.Errorf("Error run Gin router: %s \n", err))
	}

}
