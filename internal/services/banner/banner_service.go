package banner

import (
	"context"
	"github.com/daniskazan/avito-tech-banner-service/internal/dto/banners"
	"github.com/daniskazan/avito-tech-banner-service/internal/ent"
	repo "github.com/daniskazan/avito-tech-banner-service/internal/storage/repositories/banners"
)

type (
	BReader interface {
		GetBannerList() ([]ent.Banner, error)
		GetBannerByTagAndFeature(
			tagId int,
			featureId int,
			useLastRevision bool,
		) (ent.Banner, error)
	}
	BWriter interface {
		CreateBanner(
			featureId int,
			tagIds []int,
			isActive bool,
		) (int, error)
	}
)

type BannerReadService struct {
	ReadRepo repo.BannerReadRepository
}

func NewBannerReadService(repo repo.BannerReadRepository) *BannerReadService {
	return &BannerReadService{ReadRepo: repo}
}

func (s *BannerReadService) GetBannerList() ([]ent.Banner, error) {
	return nil, nil

}

func (s *BannerReadService) GetBannerByTagAndFeature(tagId int, featureId int, useLastRevision bool) (ent.Banner, error) {
	return s.ReadRepo.GetBannerByTagAndFeatureId(context.Background(), tagId, featureId)

}

type BannerWriteService struct {
	UpdateRepo repo.BannerUpdateRepository
}

func NewBannerWriteService(repo repo.BannerUpdateRepository) *BannerWriteService {
	return &BannerWriteService{UpdateRepo: repo}
}

func (s *BannerWriteService) CreateBanner(featureId int, tagIds []int, isActive bool) (int, error) {
	var dto = banners.BannerCreateDTO{
		FeatureId: featureId,
		TagIds:    tagIds,
		IsActive:  isActive,
	}
	return s.UpdateRepo.CreateBanner(context.Background(), dto)
}
