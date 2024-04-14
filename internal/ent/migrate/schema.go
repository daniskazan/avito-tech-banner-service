// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BannersColumns holds the columns for the "banners" table.
	BannersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "content", Type: field.TypeJSON},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "is_active", Type: field.TypeBool, Default: true},
	}
	// BannersTable holds the schema information for the "banners" table.
	BannersTable = &schema.Table{
		Name:       "banners",
		Columns:    BannersColumns,
		PrimaryKey: []*schema.Column{BannersColumns[0]},
	}
	// FeaturesColumns holds the columns for the "features" table.
	FeaturesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "feature_name", Type: field.TypeString, Unique: true, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// FeaturesTable holds the schema information for the "features" table.
	FeaturesTable = &schema.Table{
		Name:       "features",
		Columns:    FeaturesColumns,
		PrimaryKey: []*schema.Column{FeaturesColumns[0]},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "tag_name", Type: field.TypeString, Size: 2147483647},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}
	// BannerTagsColumns holds the columns for the "banner_tags" table.
	BannerTagsColumns = []*schema.Column{
		{Name: "banner_id", Type: field.TypeInt},
		{Name: "tag_id", Type: field.TypeInt},
	}
	// BannerTagsTable holds the schema information for the "banner_tags" table.
	BannerTagsTable = &schema.Table{
		Name:       "banner_tags",
		Columns:    BannerTagsColumns,
		PrimaryKey: []*schema.Column{BannerTagsColumns[0], BannerTagsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "banner_tags_banner_id",
				Columns:    []*schema.Column{BannerTagsColumns[0]},
				RefColumns: []*schema.Column{BannersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "banner_tags_tag_id",
				Columns:    []*schema.Column{BannerTagsColumns[1]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// FeatureBannersColumns holds the columns for the "feature_banners" table.
	FeatureBannersColumns = []*schema.Column{
		{Name: "feature_id", Type: field.TypeInt},
		{Name: "banner_id", Type: field.TypeInt},
	}
	// FeatureBannersTable holds the schema information for the "feature_banners" table.
	FeatureBannersTable = &schema.Table{
		Name:       "feature_banners",
		Columns:    FeatureBannersColumns,
		PrimaryKey: []*schema.Column{FeatureBannersColumns[0], FeatureBannersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "feature_banners_feature_id",
				Columns:    []*schema.Column{FeatureBannersColumns[0]},
				RefColumns: []*schema.Column{FeaturesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "feature_banners_banner_id",
				Columns:    []*schema.Column{FeatureBannersColumns[1]},
				RefColumns: []*schema.Column{BannersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BannersTable,
		FeaturesTable,
		TagsTable,
		BannerTagsTable,
		FeatureBannersTable,
	}
)

func init() {
	BannerTagsTable.ForeignKeys[0].RefTable = BannersTable
	BannerTagsTable.ForeignKeys[1].RefTable = TagsTable
	FeatureBannersTable.ForeignKeys[0].RefTable = FeaturesTable
	FeatureBannersTable.ForeignKeys[1].RefTable = BannersTable
}
