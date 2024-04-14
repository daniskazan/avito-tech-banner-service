package handlers

import (
	v1 "github.com/daniskazan/avito-tech-banner-service/internal/handlers/delivery/http/v1"
	"github.com/daniskazan/avito-tech-banner-service/internal/services/banner"
	repo "github.com/daniskazan/avito-tech-banner-service/internal/storage/repositories/banners"
)

func (s *Server) ConfigureRouting() {

	bannerReadRepo, bannerUpdateRepo := repo.NewBannerReadSQLRepository(s.DB), repo.NewBannerUpdateSQLRepository(s.DB)
	bannerReadService, bannerWriteService := banner.NewBannerReadService(bannerReadRepo), banner.NewBannerWriteService(bannerUpdateRepo)
	bannerHandler := v1.BannerHandler{BannerReadService: bannerReadService, BannerWriteService: bannerWriteService}

	s.Router.GET("/health", v1.Healthz())
	s.Router.POST("/banners", bannerHandler.CreateBanner())
	s.Router.GET("/banners/", bannerHandler.GetUserBanner())
}
