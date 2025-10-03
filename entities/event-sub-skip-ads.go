package entities

import "time"

type EventSubSkipAdsType string

const (
	EventSubSkipAdsUse EventSubSkipAdsType = "use"
)

type SourceSkipAdsType string

const (
	SourceAddSkipAds SourceSkipAdsType = "add"
	SourceMembership SourceSkipAdsType = "membership"
)

type EventSubSkipAds struct {
	ID            uint64              `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	UserID        string              `gorm:"size:255;not null" json:"user_id"`
	SourceSubID   uint64              `gorm:"not null" json:"source_sub_id"`                    // sub from add event or membership
	SourceSubType SourceSkipAdsType   `gorm:"type:varchar(20);not null" json:"source_sub_type"` // from add event or membership
	QuantityUsed  uint32              `gorm:"not null" json:"quantity_used"`
	AppID         string              `gorm:"size:255" json:"app_id"`
	Type          EventSubSkipAdsType `gorm:"type:varchar(20);not null" json:"type"` // use
	Description   string              `gorm:"type:text" json:"description"`
	CreatedAt     time.Time           `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time           `gorm:"autoUpdateTime" json:"updated_at"`
}
