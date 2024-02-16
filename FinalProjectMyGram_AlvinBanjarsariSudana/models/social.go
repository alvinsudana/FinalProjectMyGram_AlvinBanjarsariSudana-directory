package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Social struct {
	GormModel
	Name       string `json:"name" form:"name" valid:"required~name of social media is required"`
	Social_url string `json:"social_url" form:"social_url" valid:"required~social url is required"`
	UserID     uint
	User       *User
}

func (s *Social) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
