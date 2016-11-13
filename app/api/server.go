package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Run(message string, port int) {
	r := gin.Default()
    r.GET("/gretting", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": message,
        })
    })
    listen := fmt.Sprintf(":%d", port)
    r.Run(listen)
}