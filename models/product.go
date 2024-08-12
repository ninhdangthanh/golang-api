package models

type ProductModel struct {
	ID     uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name   string `gorm:"type:varchar(100);not null" json:"name"`
	UserID uint   `json:"user_id"`
}
