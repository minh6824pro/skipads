package entities

import "time"

type EventAddSkipAdsType string

const (
	EventAddSkipAdsPurchase EventAddSkipAdsType = "purchase"
	EventAddSkipAdsExchange EventAddSkipAdsType = "exchange"
	EventAddSkipAdsGrant    EventAddSkipAdsType = "grant"
)

type EventAddSkipAds struct {
	ID             int                 `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID         int                 `gorm:"not null" json:"user_id"`
	PackageID      *int                `json:"package_id,omitempty"`            // null if event is grant
	SourceEventID  int                 `gorm:"not null" json:"source_event_id"` // transaction id if buy, exchange id if exchange,...
	Quantity       int                 `gorm:"not null" json:"quantity"`
	QuantityUsed   int                 `gorm:"not null" json:"quantity_used"`
	MaxUsagePerDay int                 `gorm:"not null" json:"max_usage_per_day"`
	Type           EventAddSkipAdsType `gorm:"type:varchar(20);not null" json:"type"`
	Description    string              `gorm:"type:text" json:"description"`
	ExpiresAt      *time.Time          `gorm:"not null" json:"expires_at"`
	CreatedAt      time.Time           `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time           `gorm:"autoUpdateTime" json:"updated_at"`
}
