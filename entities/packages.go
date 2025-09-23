package entities

import "time"

type PackageType string

const (
	PackageTypePurchase PackageType = "purchase"
	PackageTypeExchange PackageType = "exchange"
)

type Package struct {
	ID             uint        `gorm:"primary_key" json:"id"`
	Name           string      `gorm:"name" json:"name"`
	Quantity       int32       `gorm:"quantity" json:"quantity"`
	MaxUsagePerDay uint32      `gorm:"max_usage_per_day" json:"max_usage_per_day"`
	Type           PackageType `gorm:"type" json:"type"`
	ExpiresAfter   uint32      `gorm:"expires_after" json:"expires_after"` // Expire at now() + expiresAfter * 24 * Hour
	CreatedAt      time.Time   `gorm:"created_at" json:"created_at,omitempty"`
	UpdatedAt      time.Time   `gorm:"updated_at" json:"updated_at,omitempty"`
}
