package entities

type PackageGame struct {
	PackageID string `gorm:"size:255;primary_key" json:"package_id"`
	AppID     string `gorm:"size:255;primary_key" json:"app_id"`
}
