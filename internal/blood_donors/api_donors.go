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
	"github.com/gin-gonic/gin"
)

type DonorsAPI interface {

	// CreateDonor Post /api/blood-donors/:siteId/donors
	// Registers a new donor
	CreateDonor(c *gin.Context)

	// DeleteDonor Delete /api/blood-donors/:siteId/donors/:entryId
	// Deletes specific donor registration
	DeleteDonor(c *gin.Context)

	// GetDonors Get /api/blood-donors/:siteId/donors
	// Provides the list of registered donors
	GetDonors(c *gin.Context)

	// GetDonor Get /api/blood-donors/:siteId/donors/:entryId
	// Provides details about a donor
	GetDonor(c *gin.Context)

	// UpdateDonor Put /api/blood-donors/:siteId/donors/:entryId
	// Updates specific donor
	UpdateDonor(c *gin.Context)
}
