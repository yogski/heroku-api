package main

import (
	"employee_subscription/routers"
	"os"
)

const SERVER_PORT string = "8080"

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = SERVER_PORT
	}

	r := routers.GetEngine()
	r.Run(":" + port)
}