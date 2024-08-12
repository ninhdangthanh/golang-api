package models

// Gender enumeration
type Gender string

const (
	Male   Gender = "MALE"
	Female Gender = "FEMALE"
)

type Role string

const (
	AdminRole Role = "ADMIN"
	UserRole  Role = "USER"
)

type Account string

const (
	Active   Account = "ACTIVE"
	Inactive Account = "INACTIVE"
)

type UserModel struct {
	ID       uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Email    string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password string         `gorm:"type:varchar(100);not null" json:"password"`
	Gender   Gender         `gorm:"type:varchar(6);not null" json:"gender"`
	Role     Role           `gorm:"type:varchar(5);not null" json:"role"`
	Account  Account        `gorm:"type:varchar(8);not null" json:"account"`
	Products []ProductModel `gorm:"foreignKey:UserID" json:"-"`
	Sorts    []SortModel    `gorm:"foreignKey:UserID" json:"-"`
	Tokens   []TokenModel   `gorm:"foreignKey:UserID" json:"tokens"`
}
