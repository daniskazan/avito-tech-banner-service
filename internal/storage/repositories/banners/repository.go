package banners

import (
	"context"
	"errors"
	"github.com/daniskazan/avito-tech-banner-service/internal/dto/banners"
	"github.com/daniskazan/avito-tech-banner-service/internal/ent"
	"github.com/daniskazan/avito-tech-banner-service/internal/ent/banner"
	"github.com/daniskazan/avito-tech-banner-service/internal/ent/feature"
	"github.com/daniskazan/avito-tech-banner-service/internal/ent/tag"
)

type (
	BannerReadRepository interface {
		GetBannerByTagAndFeatureId(ctx context.Context, tagId int, featureId int) (ent.Banner, error)
	}

	BannerUpdateRepository interface {
		CreateBanner(ctx context.Context, dto banners.BannerCreateDTO) (int, error)
		UpdateBanner(ctx context.Context, bannerId int) (int, error)
		DeleteBannerByID(ctx context.Context, bannerId int) (int, error)
	}
)

type (
	BannerReadSQLRepository struct {
		Client *ent.Client
	}
	BannerUpdateSQLRepository struct {
		Client *ent.Client
	}
)

func NewBannerReadSQLRepository(client *ent.Client) *BannerReadSQLRepository {
	return &BannerReadSQLRepository{Client: client}
}

func (repo *BannerReadSQLRepository) GetBannerByTagAndFeatureId(ctx context.Context, tagId int, featureId int) (ent.Banner, error) {

	return ent.Banner{}, nil
}

func NewBannerUpdateSQLRepository(client *ent.Client) *BannerUpdateSQLRepository {
	return &BannerUpdateSQLRepository{Client: client}
}

func (repo *BannerUpdateSQLRepository) CreateBanner(ctx context.Context, dto banners.BannerCreateDTO) (int, error) {
	query := repo.Client.Banner.Query().
		Where(
			banner.HasFeatureWith(
				feature.ID(dto.FeatureId),
			),
		).
		Where(
			banner.HasTagsWith(
				tag.IDIn(dto.TagIds...),
			),
		)

	// Проверка наличия совпадений.
	exists, _ := query.Exist(context.Background())
	if exists {
		return 0, errors.New("Баннер уже существует")
	}

	b, err := repo.Client.Banner.Create().SetContent(dto.Content).SetIsActive(dto.IsActive).AddFeatureIDs(dto.FeatureId).AddTagIDs(dto.TagIds...).Save(ctx)
	if err != nil {
		return 0, err
	}
	return b.ID, nil
}
func (repo *BannerUpdateSQLRepository) UpdateBanner(ctx context.Context, bannerId int) (int, error) {
	return bannerId, nil
}
func (repo *BannerUpdateSQLRepository) DeleteBannerByID(ctx context.Context, bannerId int) (int, error) {
	return bannerId, nil
}
