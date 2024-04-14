package v1

import (
	"github.com/daniskazan/avito-tech-banner-service/internal/services/banner"
	"github.com/labstack/echo/v4"
	"net/http"
)

type BannerHandler struct {
	BannerReadService  banner.BReader
	BannerWriteService banner.BWriter
}

func NewBannerHandler(rs banner.BReader, ws banner.BWriter) *BannerHandler {
	return &BannerHandler{BannerReadService: rs, BannerWriteService: ws}
}
func (h *BannerHandler) GetUserBanner() echo.HandlerFunc {
	type (
		Request struct {
			FeatureID       int  `query:"featureId"`
			TagID           int  `query:"tagId"`
			UseLastRevision bool `query:"useLastRevision"`
		}
		Response struct {
			BannerContent map[string]any `json:"bannerContent"`
		}
	)
	var request Request
	return func(c echo.Context) error {
		err := c.Bind(&request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Некорректные данные"})
		}

		b, err := h.BannerReadService.GetBannerByTagAndFeature(request.TagID, request.FeatureID, request.UseLastRevision)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, Response{
			BannerContent: b.Content,
		})

	}
}
func (h *BannerHandler) CreateBanner() echo.HandlerFunc {
	type (
		Request struct {
			TagIds    []int          `json:"tagIds" `
			FeatureId int            `json:"featureId"`
			IsActive  bool           `json:"isActive"`
			Content   map[string]any `json:"content"`
		}
		Response struct {
			BannerID int `json:"bannerId"`
		}
	)

	var request Request
	return func(c echo.Context) error {
		err := c.Bind(&request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Некорректные данные"})
		}

		bannerId, err := h.BannerWriteService.CreateBanner(request.FeatureId, request.TagIds, request.IsActive)
		if err != nil {
			// TODO case type of error return appropriate status
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusCreated, Response{
			bannerId,
		})

	}
}
