package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Gretting(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Welcome . . . to API command line",
	})
}