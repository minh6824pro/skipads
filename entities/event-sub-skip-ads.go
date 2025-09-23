package entities

import "time"

type EventSubSkipAdsType string

const (
	EventSubSkipAdsUse EventSubSkipAdsType = "use"
)

type EventSubSkipAds struct {
	ID           int                 `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID       int                 `gorm:"not null" json:"user_id"`
	SourceSubID  int                 `gorm:"not null" json:"source_sub_id"` // sub from add event or membership
	QuantityUsed int                 `gorm:"not null" json:"quantity_used"`
	Type         EventSubSkipAdsType `gorm:"type:varchar(20);not null" json:"type"`
	Description  string              `gorm:"type:text" json:"description"`
	CreatedAt    time.Time           `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time           `gorm:"autoUpdateTime" json:"updated_at"`
}
