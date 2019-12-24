package router

import (
	"github.com/gin-gonic/gin"
	"github.com/moeen/atlas-dns/handlers"
)

// Router returns the gin router with all of our defined routes
func Router() *gin.Engine {
	r := gin.Default()

	// At the moment we only have v1 of our API
	v1 := r.Group("/v1")
	{
		// dns route is responsible for calculating location from given data
		v1.POST("/dns", handlers.DnsHandler)
	}

	return r
}
