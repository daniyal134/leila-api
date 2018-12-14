package main

import (
	"github.com/leilaApp/leila-api/routers"
	"os"
)



const ENV_RUN_PORT = "PORT"
var RUN_PORT string = "8080"

func init() {
	getEnvPort()
}

func getEnvPort() {
	port := os.Getenv(ENV_RUN_PORT)

	if len(port) > 0 {
		RUN_PORT = port
	}
}

func main() {

	router := routers.InitRoutes()
	router.Run(":" + RUN_PORT)
}
