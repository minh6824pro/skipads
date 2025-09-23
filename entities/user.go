package entities

type User struct {
	UserID int    `gorm:"not null" json:"user_id"`
	Name   string `gorm:"not null" json:"name"`
}
