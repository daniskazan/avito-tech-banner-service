package entities

type (
	BannerID  int
	FeatureID int
	TagID     int
)

type BannerEntity struct {
	BannerId BannerID
	Feature  FeatureID
	Tags     []TagID
}
