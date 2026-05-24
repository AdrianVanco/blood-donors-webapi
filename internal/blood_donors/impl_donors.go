package blood_donors

import (
	"net/http"

	"slices"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type implDonorsAPI struct {
	logger zerolog.Logger
}

func NewDonorsApi() DonorsAPI {
	return &implDonorsAPI{logger: log.With().Str("component", "blood-donors").Logger()}
}

func (o implDonorsAPI) CreateDonor(c *gin.Context) {
	updateSiteFunc(c, func(c *gin.Context, site *DonationSite) (*DonationSite, interface{}, int) {
		logger := o.logger.With().
			Str("method", "CreateDonor").
			Str("siteId", site.Id).
			Logger()
		var donor Donor

		if err := c.ShouldBindJSON(&donor); err != nil {
			logger.Error().Err(err).Msg("Failed to bind JSON")
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid request body",
				"error":   err.Error(),
			}, http.StatusBadRequest
		}

		if donor.DonorId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Donor ID is required",
			}, http.StatusBadRequest
		}

		if donor.Id == "" || donor.Id == "@new" {
			donor.Id = uuid.NewString()
		}

		// registračné číslo darcu musí byť unikátne v rámci miesta
		conflictIndx := slices.IndexFunc(site.Donors, func(d Donor) bool {
			return donor.Id == d.Id || donor.DonorId == d.DonorId
		})
		if conflictIndx >= 0 {
			return nil, gin.H{
				"status":  http.StatusConflict,
				"message": "Donor already exists",
			}, http.StatusConflict
		}

		site.Donors = append(site.Donors, donor)
		logger.Info().Str("entry-id", donor.Id).Msg("Successfully created donor")
		return site, donor, http.StatusOK
	})
}

func (o implDonorsAPI) DeleteDonor(c *gin.Context) {
	updateSiteFunc(c, func(c *gin.Context, site *DonationSite) (*DonationSite, interface{}, int) {
		entryId := c.Param("entryId")

		if entryId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Entry ID is required",
			}, http.StatusBadRequest
		}

		indx := slices.IndexFunc(site.Donors, func(d Donor) bool {
			return entryId == d.Id
		})

		if indx < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Donor not found",
			}, http.StatusNotFound
		}

		site.Donors = append(site.Donors[:indx], site.Donors[indx+1:]...)
		return site, nil, http.StatusNoContent
	})
}

func (o implDonorsAPI) GetDonors(c *gin.Context) {
	updateSiteFunc(c, func(c *gin.Context, site *DonationSite) (*DonationSite, interface{}, int) {
		result := site.Donors
		if result == nil {
			result = []Donor{}
		}
		// return nil site - no need to update it in db
		return nil, result, http.StatusOK
	})
}

func (o implDonorsAPI) GetDonor(c *gin.Context) {
	updateSiteFunc(c, func(c *gin.Context, site *DonationSite) (*DonationSite, interface{}, int) {
		entryId := c.Param("entryId")

		if entryId == "" {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Entry ID is required",
			}, http.StatusBadRequest
		}

		indx := slices.IndexFunc(site.Donors, func(d Donor) bool {
			return entryId == d.Id
		})

		if indx < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Donor not found",
			}, http.StatusNotFound
		}

		// return nil site - no need to update it in db
		return nil, site.Donors[indx], http.StatusOK
	})
}

func (o implDonorsAPI) UpdateDonor(c *gin.Context) {
	updateSiteFunc(c, func(c *gin.Context, site *DonationSite) (*DonationSite, interface{}, int) {
		var donor Donor

		if err := c.ShouldBindJSON(&donor); err != nil {
			return nil, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Invalid request body",
				"error":   err.Error(),
			}, http.StatusBadRequest
		}

		entryId := c.Param("entryId")

		indx := slices.IndexFunc(site.Donors, func(d Donor) bool {
			return entryId == d.Id
		})

		if indx < 0 {
			return nil, gin.H{
				"status":  http.StatusNotFound,
				"message": "Donor not found",
			}, http.StatusNotFound
		}

		// id a registračné číslo sú nemenné - zachováme pôvodné hodnoty
		donor.Id = site.Donors[indx].Id
		if donor.DonorId == "" {
			donor.DonorId = site.Donors[indx].DonorId
		}

		site.Donors[indx] = donor
		return site, site.Donors[indx], http.StatusOK
	})
}
