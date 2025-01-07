package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	"alive/alive/routes"
)

var (
	defaultListenHost = "::"
	defaultListenPort = 4444
)

func getListenHost() string {
	host := os.Getenv("LISTEN_HOST")
	if host == "" {
		host = defaultListenHost
	}
	return host
}

func getListenPort() int {
	portStr := os.Getenv("LISTEN_PORT")
	if portStr == "" {
		portStr = strconv.Itoa(defaultListenPort)
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(fmt.Sprintf("Invalid port number: %s", err))
	}

	return port
}

func formatAddress(host string, port int) string {
	if host == defaultListenHost {
		return fmt.Sprintf("[%s]:%d", host, port)
	}
	return fmt.Sprintf("%s:%d", host, port)
}

func main() {
	if gin.Mode() == gin.ReleaseMode {
		gin.DefaultWriter = io.Discard
	} else {
		gin.DefaultWriter = os.Stdout
	}

	router := gin.Default()

	router.GET("/healthz", routes.Health)
	router.GET("/liveness", routes.Liveness)

	host := getListenHost()
	port := getListenPort()
	addr := formatAddress(host, port)
	fmt.Println("Listening on", addr)

	err := router.Run(addr)
	if err != nil {
		panic(fmt.Sprintf("Error starting server: %s", err))
	}
}
