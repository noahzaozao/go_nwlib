package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Uuid     string `gorm:"TYPE:VARCHAR(36);UNIQUE_INDEX;NOT NULL"`
	Mobile   string `gorm:"TYPE:VARCHAR(255);"`
	Password string `gorm:"TYPE:VARCHAR(255);"`
	Username string `gorm:"TYPE:VARCHAR(255);"`
	Email    string `gorm:"TYPE:VARCHAR(255);"`
	Token    string `gorm:"TYPE:VARCHAR(255);"`
}

type WechatMappUser struct {
	gorm.Model
	UserID    uint
	User      User   `gorm:"FOREIGNKEY:ID;ASSOCIATION_FOREIGNKEY:UserID"`
	Nickname  string `gorm:"TYPE:VARCHAR(255);"`
	Gender    string `gorm:"TYPE:VARCHAR(255);"`
	City      string `gorm:"TYPE:VARCHAR(255);"`
	Country   string `gorm:"TYPE:VARCHAR(255);"`
	Province  string `gorm:"TYPE:VARCHAR(255);"`
	AvatarUrl string `gorm:"TYPE:VARCHAR(255);"`
	OpenId    string `gorm:"TYPE:VARCHAR(255);"`
	UnionId   string `gorm:"TYPE:VARCHAR(255);"`
	AppID	  string `gorm:"TYPE:VARCHAR(255);"`
}
