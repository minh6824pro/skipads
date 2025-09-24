package entities

import "time"

type PackageType string

const (
	PackageTypePurchase PackageType = "purchase"
	PackageTypeExchange PackageType = "exchange"
)

type Package struct {
	ID           uint32      `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name         string      `gorm:"name" json:"name"`
	Quantity     uint32      `gorm:"quantity" json:"quantity"`
	Type         PackageType `gorm:"type:varchar(20)" json:"type"`
	ExpiresAfter uint32      `gorm:"expires_after" json:"expires_after"` // Expire at now() + expiresAfter * 24 * Hour
	CreatedAt    time.Time   `gorm:"created_at" json:"created_at,omitempty"`
	UpdatedAt    time.Time   `gorm:"updated_at" json:"updated_at,omitempty"`
}
