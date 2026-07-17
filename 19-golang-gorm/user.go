package golang_gorm

import (
	"time"
)

// User => users, otomatis (default convention) di mapping menjadi table users di database
type User struct {
	ID          string    `gorm:"primary_key;column:id;<-:create"`
	Password    string    `gorm:"column:password"`
	Name        string    `gorm:"column:name"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Information string    `gorm:"-"`
}

// Override nama table
func (u *User) TableName() string {
	return "users" // nama table
}
