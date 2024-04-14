package banners

import (
	e "github.com/daniskazan/avito-tech-banner-service/internal/entities"
	"gorm.io/gorm"
)

type (
	BannerReadRepository interface {
		GetBannerByTagAndFeatureId(tag e.TagID, feature e.FeatureID) (e.BannerEntity, error)
	}

	BannerUpdateRepository interface {
		CreateBanner(banner e.BannerEntity) (e.BannerID, error)
		UpdateBanner(banner e.BannerID) (e.BannerID, error)
		DeleteBannerByID(bannerId e.BannerID) (e.BannerID, error)
	}
)

type (
	BannerReadSQLRepository struct {
		DB *gorm.DB
	}
	BannerUpdateSQLRepository struct {
		DB *gorm.DB
	}
)

func NewBannerReadSQLRepository(db *gorm.DB) *BannerReadSQLRepository {
	return &BannerReadSQLRepository{DB: db}
}

func (repo *BannerReadSQLRepository) GetBannerByTagAndFeatureId(tag e.TagID, feature e.FeatureID) (e.BannerEntity, error) {
	var banner e.BannerEntity
	repo.DB.Raw("select * from banner_entities b join feature_entities fe on fe.id = b.feature_id join public.banner_tags bt on b.id = bt.banner_entity_id where fe.id = ? and bt.tag_entity_id = ?", feature, tag).Scan(&banner)
	return banner, nil
}

func NewBannerUpdateSQLRepository(db *gorm.DB) *BannerUpdateSQLRepository {
	return &BannerUpdateSQLRepository{DB: db}
}

func (repo *BannerUpdateSQLRepository) CreateBanner(banner e.BannerEntity) (e.BannerID, error) {
	return e.BannerID(1), nil
}
func (repo *BannerUpdateSQLRepository) UpdateBanner(banner e.BannerID) (e.BannerID, error) {
	return e.BannerID(1), nil
}
func (repo *BannerUpdateSQLRepository) DeleteBannerByID(bannerId e.BannerID) (e.BannerID, error) {
	return e.BannerID(1), nil
}
