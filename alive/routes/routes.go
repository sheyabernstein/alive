// routes.go

package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Liveness(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
