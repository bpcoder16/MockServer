package v1

import (
	"MockServer/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonApiPort(port string) gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.FullPath()[1:]
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.String(http.StatusOK, global.MockServerDict[port][uri].Response)
	}
}
