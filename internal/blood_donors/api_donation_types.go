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

type DonationTypesAPI interface {

	// GetDonationTypes Get /api/blood-donors/:siteId/donation-types
	// Provides the list of donation types offered at the donation site
	GetDonationTypes(c *gin.Context)
}
