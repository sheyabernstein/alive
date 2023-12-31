package main

import (
	"os"
	"fmt"
	"strconv"
	"net/http"

	"github.com/gin-gonic/gin"
)


// getLiveness responds with HTTP 200 OK
func getLiveness(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func main() {
	router := gin.Default()
	
	router.GET("/liveness", getLiveness)
	
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "4444"
	}
	
	// convert the port string to an integer
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(fmt.Sprintf("Invalid port number: ", err))
	}

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	fmt.Println("Listening on", addr)
	router.Run(addr)
}