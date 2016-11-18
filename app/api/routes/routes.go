package routes

import (
	"github.com/gin-gonic/gin"

	c "app/api/controllers"
)

// Load routes function
func Load_routes(router_group *gin.RouterGroup) {
	// Gretting routes
	router_group.GET("/gretting", c.Gretting)
}