package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"docker-alive/alive/routes"
)

var defaultPort int = 4444

func getPort() int {
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = strconv.Itoa(defaultPort)
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(fmt.Sprintf("Invalid port number: %s", err))
	}

	return port
}

func main() {
	router := gin.Default()

	router.GET("/healthz", routes.Health)
	router.GET("/liveness", routes.Liveness)

	port := getPort()
	addr := fmt.Sprintf("0.0.0.0:%d", port)
	fmt.Println("Listening on", addr)

	err := router.Run(addr)
	if err != nil {
		panic(fmt.Sprintf("Error starting server: %s", err))
	}
}
