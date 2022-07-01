package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name     string `json:"name"`
	ThumbUrl string `json:"thumb_url"`
}

type Category struct {
	gorm.Model
	Name     string `json:"name"`
	ThumbUrl string `json:"thumb_url"`
}

type Video struct {
	gorm.Model
	Name       string   `json:"name"`
	Url        string   `json:"url"`
	CategoryID uint     `json:"-"`
	Category   Category `json:"category"`
	Tags       []Tag    `gorm:"many2many:video_tags" json:"tags"`
}
