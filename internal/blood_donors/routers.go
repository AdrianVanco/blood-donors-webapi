/*
 * Blood Donors Api
 *
 * Blood donor registration and management for Web-In-Cloud system
 *
 * API version: 1.0.0
 * Contact: vancoa@stuba.sk
 */

package blood_donors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

// NewRouter returns a new router.
func NewRouter(handleFunctions ApiHandleFunctions) *gin.Engine {
	return NewRouterWithGinEngine(gin.Default(), handleFunctions)
}

// NewRouterWithGinEngine add routes to existing gin engine.
func NewRouterWithGinEngine(router *gin.Engine, handleFunctions ApiHandleFunctions) *gin.Engine {
	for _, route := range getRoutes(handleFunctions) {
		if route.HandlerFunc == nil {
			route.HandlerFunc = DefaultHandleFunc
		}
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			router.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(route.Pattern, route.HandlerFunc)
		}
	}

	return router
}

// DefaultHandleFunc Default handler for not yet implemented routes
func DefaultHandleFunc(c *gin.Context) {
	c.String(http.StatusNotImplemented, "501 not implemented")
}

type ApiHandleFunctions struct {

	// Routes for the DonationTypesAPI part of the API
	DonationTypesAPI DonationTypesAPI
	// Routes for the DonorsAPI part of the API
	DonorsAPI DonorsAPI
	// Routes for the DonationSitesAPI part of the API
	DonationSitesAPI DonationSitesAPI
}

func getRoutes(handleFunctions ApiHandleFunctions) []Route {
	return []Route{
		{
			"GetDonationTypes",
			http.MethodGet,
			"/api/blood-donors/:siteId/donation-types",
			handleFunctions.DonationTypesAPI.GetDonationTypes,
		},
		{
			"CreateDonor",
			http.MethodPost,
			"/api/blood-donors/:siteId/donors",
			handleFunctions.DonorsAPI.CreateDonor,
		},
		{
			"DeleteDonor",
			http.MethodDelete,
			"/api/blood-donors/:siteId/donors/:entryId",
			handleFunctions.DonorsAPI.DeleteDonor,
		},
		{
			"GetDonors",
			http.MethodGet,
			"/api/blood-donors/:siteId/donors",
			handleFunctions.DonorsAPI.GetDonors,
		},
		{
			"GetDonor",
			http.MethodGet,
			"/api/blood-donors/:siteId/donors/:entryId",
			handleFunctions.DonorsAPI.GetDonor,
		},
		{
			"UpdateDonor",
			http.MethodPut,
			"/api/blood-donors/:siteId/donors/:entryId",
			handleFunctions.DonorsAPI.UpdateDonor,
		},
		{
			"CreateDonationSite",
			http.MethodPost,
			"/api/blood-donors",
			handleFunctions.DonationSitesAPI.CreateDonationSite,
		},
		{
			"DeleteDonationSite",
			http.MethodDelete,
			"/api/blood-donors/:siteId",
			handleFunctions.DonationSitesAPI.DeleteDonationSite,
		},
	}
}
