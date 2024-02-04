package main

import (
	"github.com/gin-gonic/gin"

	router "myapp_scenario_api/internal/app/router"
)

func init() {
	// set server mode
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	r := gin.Default()

	// set trusted proxies
	r.SetTrustedProxies([]string{"localhost", "https://myapp-bff.onrender.com"})

	// set routes
	router.SetupRouter(r)

	r.Run(":8000")
}
