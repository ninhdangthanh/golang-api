package models

type SortModel struct {
	ID     uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Type   string `gorm:"type:varchar(100);not null" json:"type"`
	UserID uint   `json:"user_id"`
}
