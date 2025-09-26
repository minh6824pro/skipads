package entities

import "time"

type EventAddSkipAdsArchive struct {
	ID            uint32              `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        uint32              `gorm:"not null" json:"user_id"`
	PackageID     *uint32             `json:"package_id,omitempty"`            // null if event is grant
	SourceEventID uint32              `gorm:"not null" json:"source_event_id"` // transaction id if buy, exchange id if exchange,...
	Quantity      uint32              `gorm:"not null" json:"quantity"`
	QuantityUsed  uint32              `gorm:"not null" json:"quantity_used"`
	Type          EventAddSkipAdsType `gorm:"type:varchar(20);not null" json:"type"`
	Description   string              `gorm:"type:text" json:"description"`
	ExpiresAt     time.Time           `gorm:"not null" json:"expires_at"`
	CreatedAt     time.Time           `json:"created_at"`
	UpdatedAt     time.Time           `json:"updated_at"`
}
