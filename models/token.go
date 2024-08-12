package models

type TokenModel struct {
	ID     uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Value  string `gorm:"type:varchar(100);not null" json:"value"`
	UserID uint   `json:"user_id"`
}
