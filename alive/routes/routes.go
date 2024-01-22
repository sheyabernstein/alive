// routes.go

package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Health(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func Liveness(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
