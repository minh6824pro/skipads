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

// index idx_user_expire(user_id,priority,expires_at)
type EventAddSkipAds struct {
	ID            uint32              `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID        uint32              `gorm:"not null;index:idx_user_expire" json:"user_id"`
	PackageID     *uint32             `json:"package_id,omitempty"`            // null if event is grant
	SourceEventID uint32              `gorm:"not null" json:"source_event_id"` // transaction id if buy, exchange id if exchange,...
	Quantity      uint32              `gorm:"not null" json:"quantity"`
	QuantityUsed  uint32              `gorm:"not null" json:"quantity_used"`
	Type          EventAddSkipAdsType `gorm:"type:varchar(20);not null" json:"type"`
	Description   string              `gorm:"type:text" json:"description"`
	Priority      uint32              `gorm:"not null;index:idx_user_expire" json:"priority"`
	ExpiresAt     time.Time           `gorm:"not null;index:idx_user_expire" json:"expires_at"`
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
