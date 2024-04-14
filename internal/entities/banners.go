package entities

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type (
	BannerID  int
	FeatureID int
	TagID     int
)

type BannerEntity struct {
	gorm.Model
	Content   datatypes.JSON `json:"content"` // JSON-документ неопределенной структуры
	IsActive  bool           `json:"is_active"`
	FeatureID int            `json:"feature_id"` // ID фичи
	Feature   FeatureEntity  `gorm:"foreignKey:FeatureID"`
	Tags      []TagEntity    `gorm:"many2many:banner_tags;"`
}

type TagEntity struct {
	gorm.Model
	Name    string         `json:"name"`
	Banners []BannerEntity `gorm:"many2many:banner_tags;"`
}

type FeatureEntity struct {
	gorm.Model
	Name    string         `json:"name"`
	Banners []BannerEntity `gorm:"foreignKey:FeatureID"`
}
