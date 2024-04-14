package v1

import (
	e "github.com/daniskazan/avito-tech-banner-service/internal/entities"
	"github.com/daniskazan/avito-tech-banner-service/internal/services/banner"
	"github.com/labstack/echo/v4"
	"net/http"
)

type BannerHandler struct {
	BannerReadService  banner.BReader
	BannerWriteService banner.BWriter
}

func (h *BannerHandler) CreateBanner() echo.HandlerFunc {
	type (
		Request struct {
			TagIds    []e.TagID   `json:"tagIds" `
			FeatureId e.FeatureID `json:"featureId"`
			IsActive  bool
		}
		Response struct {
			BannerID e.BannerID `json:"bannerId"`
		}
		CreateBannerDTO struct {
			TagIds    []e.TagID
			FeatureId e.FeatureID
			IsActive  bool
		}
	)

	var request Request
	return func(c echo.Context) error {
		err := c.Bind(&request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Некорректные данные"})
		}

		var dto = CreateBannerDTO{
			FeatureId: request.FeatureId,
			TagIds:    request.TagIds,
			IsActive:  request.IsActive,
		}
		bannerId, err := h.BannerWriteService.CreateBanner(dto.FeatureId, dto.TagIds, dto.IsActive)
		if err != nil {
			// TODO case type of error return appropriate status
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusCreated, Response{
			bannerId,
		})

	}
}
