package entities

import "time"

type EventAddSkipAdsType string

const (
	EventAddSkipAdsPurchase EventAddSkipAdsType = "purchase"
	EventAddSkipAdsExchange EventAddSkipAdsType = "exchange"
	EventAddSkipAdsGrant    EventAddSkipAdsType = "grant"
)

var priorityMap = map[EventAddSkipAdsType]uint32{
	EventAddSkipAdsExchange: 1,
	EventAddSkipAdsGrant:    2,
	EventAddSkipAdsPurchase: 3,
}

type EventAddSkipAds struct {
	ID            uint64              `gorm:"primaryKey;AUTO_INCREMENT" json:"id"`
	UserID        string              `gorm:"size:255;not null" json:"user_id"`
	PackageID     *string             `json:"size:255;package_id,omitempty"`            // null if event is grant
	SourceEventID string              `gorm:"size:255;not null" json:"source_event_id"` // transaction id if buy, exchange id if exchange,...
	Quantity      uint32              `gorm:"not null" json:"quantity"`
	QuantityUsed  uint32              `gorm:"not null" json:"quantity_used"`
	Type          EventAddSkipAdsType `gorm:"type:varchar(20);not null" json:"type"`
	Description   string              `gorm:"size:255;type:text" json:"description"`
	Priority      uint32              `gorm:"not null" json:"priority"`
	ExpiresAt     time.Time           `gorm:"not null" json:"expires_at"`
	CreatedAt     time.Time           `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time           `gorm:"autoUpdateTime" json:"updated_at"`
}

func (e *EventAddSkipAds) ConvertToEventAddSkipAdsArchive() EventAddSkipAdsArchive {
	event := EventAddSkipAdsArchive{
		ID:            e.ID,
		UserID:        e.UserID,
		PackageID:     e.PackageID,
		SourceEventID: e.SourceEventID,
		Quantity:      e.Quantity,
		QuantityUsed:  e.QuantityUsed,
		Type:          e.Type,
		Description:   e.Description,
		ExpiresAt:     e.ExpiresAt,
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
	}
	return event
}

func (e *EventAddSkipAds) SetPriority() {
	if p, ok := priorityMap[e.Type]; ok {
		e.Priority = p
	} else {
		e.Priority = 3 // default
	}
}
