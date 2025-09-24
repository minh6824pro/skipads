package entities

type User struct {
	UserID uint32 `gorm:"primary_key;AUTO_INCREMENT" json:"user_id"`
	Name   string `gorm:"not null" json:"name"`
}
