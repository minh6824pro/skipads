package entities

type PackageType string

const (
	PackageTypePurchase PackageType = "purchase"
	PackageTypeExchange PackageType = "exchange"
)

type Package struct {
	ID           string      `gorm:"size:255;primary_key" json:"id"`
	Name         string      `gorm:"size:255;name" json:"name"`
	Quantity     uint32      `gorm:"quantity" json:"quantity"`
	Type         PackageType `gorm:"type:varchar(20)" json:"type"`
	ExpiresAfter uint32      `gorm:"expires_after" json:"expires_after"` // Expire at now() + expiresAfter * 24 * Hour
}
