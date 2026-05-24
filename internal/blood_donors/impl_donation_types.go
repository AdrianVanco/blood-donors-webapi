package blood_donors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type implDonationTypesAPI struct {
}

func NewDonationTypesApi() DonationTypesAPI {
	return &implDonationTypesAPI{}
}

func (o implDonationTypesAPI) GetDonationTypes(c *gin.Context) {
	updateSiteFunc(c, func(
		c *gin.Context,
		site *DonationSite,
	) (updatedSite *DonationSite, responseContent interface{}, status int) {
		result := site.PredefinedDonationTypes
		if result == nil {
			result = []DonationType{}
		}
		return nil, result, http.StatusOK
	})
}
