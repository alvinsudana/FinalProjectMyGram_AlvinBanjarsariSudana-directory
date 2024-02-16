package models

import (
	"MyGram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Email    string    `gorm:"not null; uniqueIndex" json:"email" form:"email" valid:"required~ Your email is required,email~Invalid email format"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Username string    `gorm:"not null" json:"username" form:"username" valid:"required~Your username is required"`
	Age      uint      `gorm:"not null" json:"age" form:"age" valid:"required~Your age is required,range(8|1000)~The minimum Age is 8"`
	Photos   []Photo   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photo"`
	Socials  []Social  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social"`
	Comments []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comment"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
