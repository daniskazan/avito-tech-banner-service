package banners

type BannerCreateDTO struct {
	TagIds    []int
	FeatureId int
	IsActive  bool
	Content   map[string]any
}
