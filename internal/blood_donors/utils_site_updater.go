package blood_donors

import (
	"net/http"

	"github.com/AdrianVanco/blood-donors-webapi/internal/db_service"
	"github.com/gin-gonic/gin"
)

type siteUpdater = func(
	ctx *gin.Context,
	site *DonationSite,
) (updatedSite *DonationSite, responseContent interface{}, status int)

// updateSiteFunc načíta odberné miesto z DB, spustí na ňom updater a prípadne uloží zmenu
func updateSiteFunc(ctx *gin.Context, updater siteUpdater) {
	value, exists := ctx.Get("db_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service not found",
				"error":   "db_service not found",
			})
		return
	}

	db, ok := value.(db_service.DbService[DonationSite])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db_service context is not of type db_service.DbService",
				"error":   "cannot cast db_service context to db_service.DbService",
			})
		return
	}

	siteId := ctx.Param("siteId")

	site, err := db.FindDocument(ctx, siteId)

	switch err {
	case nil:
		// continue
	case db_service.ErrNotFound:
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Donation site not found",
				"error":   err.Error(),
			},
		)
		return
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to load donation site from database",
				"error":   err.Error(),
			})
		return
	}

	updatedSite, responseObject, status := updater(ctx, site)

	if updatedSite != nil {
		err = db.UpdateDocument(ctx, siteId, updatedSite)
	} else {
		err = nil // redundant but for clarity
	}

	switch err {
	case nil:
		if responseObject != nil {
			ctx.JSON(status, responseObject)
		} else {
			ctx.AbortWithStatus(status)
		}
	case db_service.ErrNotFound:
		ctx.JSON(
			http.StatusNotFound,
			gin.H{
				"status":  "Not Found",
				"message": "Donation site was deleted while processing the request",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to update donation site in database",
				"error":   err.Error(),
			})
	}
}
