package main

import (
	"fmt"

	"acgfate/conf"
	"acgfate/router"
)

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
