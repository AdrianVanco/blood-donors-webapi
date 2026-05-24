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

type DonationSitesAPI interface {

	// CreateDonationSite Post /api/blood-donors
	// Saves new donation site definition
	CreateDonationSite(c *gin.Context)

	// DeleteDonationSite Delete /api/blood-donors/:siteId
	// Deletes specific donation site
	DeleteDonationSite(c *gin.Context)
}
