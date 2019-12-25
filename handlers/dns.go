package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/moeen/atlas-dns/resources"
	"github.com/moeen/atlas-dns/services"
	"net/http"
)

// DnsHandler is responsible for POST requests from /v1/dns
func DnsHandler(ctx *gin.Context) {
	var r resources.DnsRequest
	// Trying to unmarshal the incoming POST request body to DnsRequest struct
	// It also validates the Json
	if err := ctx.ShouldBindJSON(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, resources.ErrorResponse{Message: "bad request body"})
		return
	}

	// Calculate the location with given coordinates and velocity
	l := services.FindLocation(*r.X, *r.Y, *r.Z, *r.Vel)

	// Create the response object and returns it as json with 200 status code
	res := resources.DnsResponse{Loc: l}
	ctx.JSON(http.StatusOK, res)
}
