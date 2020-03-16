package user

import "github.com/jinzhu/gorm"

type Members struct {
	gorm.Model
	Username string `gorm:"type:varchar(100)"`
	Email    string `gorm:"type:varchar(100)"`
	Password string `gorm:"type:varchar(100)"`
}
