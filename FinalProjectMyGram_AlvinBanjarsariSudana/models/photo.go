package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title     string `json:"title" form:"title" valid:"required~Title of your photo is required"`
	Photo_url string `json:"photo_url" form:"photo_url" valid:"required~Photo url is required"`
	Caption   string `json:"caption" form:"caption"`
	UserID    uint
	User      *User
	Comments  []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
