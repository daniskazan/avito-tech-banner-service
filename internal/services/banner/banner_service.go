package banner

import (
	e "github.com/daniskazan/avito-tech-banner-service/internal/entities"
	repo "github.com/daniskazan/avito-tech-banner-service/internal/storage/repositories/banners"
	"math/rand"
)

type (
	BReader interface {
		GetBannerList() ([]e.BannerEntity, error)
		GetBannerByTagAndFeature(
			tag e.TagID,
			feature e.FeatureID,
			useLastRevision bool,
		) (e.BannerEntity, error)
	}
	BWriter interface {
		CreateBanner(
			featureId e.FeatureID,
			tagIds []e.TagID,
			isActive bool,
		) (e.BannerID, error)
	}
)

type BannerReadService struct {
	ReadRepo repo.BannerReadRepository
}

func NewBannerReadService(repo repo.BannerReadRepository) *BannerReadService {
	return &BannerReadService{ReadRepo: repo}
}

func (s *BannerReadService) GetBannerList() ([]e.BannerEntity, error) {
	return nil, nil

}

func (s *BannerReadService) GetBannerByTagAndFeature(tag e.TagID, feature e.FeatureID, useLastRevision bool) (e.BannerEntity, error) {
	return s.ReadRepo.GetBannerByTagAndFeatureId(tag, feature)

}

type BannerWriteService struct {
	UpdateRepo repo.BannerUpdateRepository
}

func NewBannerWriteService(repo repo.BannerUpdateRepository) *BannerWriteService {
	return &BannerWriteService{UpdateRepo: repo}
}

func (s *BannerWriteService) CreateBanner(featureId e.FeatureID, tagIds []e.TagID, isActive bool) (e.BannerID, error) {
	return e.BannerID(rand.Int()), nil
}
