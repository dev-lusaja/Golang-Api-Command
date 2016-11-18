package api

import (
	"fmt"

	"github.com/gin-gonic/gin"

    routes "app/api/routes"
)

func Run(port int, version string) {
	router := gin.Default()
    router_group := router.Group(version)
    // Load Routes
    routes.Load_routes(router_group)

    listen_port := fmt.Sprintf(":%d", port)
    router.Run(listen_port)
}