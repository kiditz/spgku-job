package entity

import "time"

// Campaign is model for database campaigns
type Campaign struct {
	ID           uint           `gorm:"primary_key" json:"id"`
	Title        string         `json:"title" sql:"index" gorm:"type:varchar(60);index:campaign_title_idx;not null" validate:"required"`
	Method       string         `json:"method" gorm:"type:varchar(60)"  validate:"required"`
	Detail       string         `json:"detail" validate:"required" gorm:"not null"`
	Criteria     string         `json:"criteria" validate:"required" gorm:"not null"`
	Task         string         `json:"task" validate:"required" gorm:"not null"`
	SocialMedias *[]SocialMedia `json:"social_media" validate:"required" gorm:"many2many:social_media_list;"`
	// Location      *Location      `json:"location,omitempty" gorm:"foreignkey:LocationID"`
	LocationID    uint   `json:"location_id"`
	SampleProduct string `json:"sample_product"`
	// Images        *[]Image   `json:"images"`
	StartDate   *time.Time `gorm:"not null" json:"start_date" validate:"required"`
	EndDate     *time.Time `gorm:"not null" json:"end_date" validate:"gtefield=StartDate,required"`
	StartPrice  float64    `json:"start_price" gorm:"not null" validate:"gte=100000.0,required"`
	EndPrice    float64    `json:"end_price" gorm:"not null" validate:"gtefield=StartPrice,required"`
	StaffAmount uint       `json:"staff_amount"`
	Model
}
