package models

import (
	"apna-restaurant/utils"
	"strings"

	"gorm.io/gorm"
)

func (user *User) ValidateReqBody(db *gorm.DB, flag string) (string, bool) {
	//better way to write checks
	if len(strings.TrimSpace(user.Email)) == 0 {
		return "Email required", false
	} else if !utils.IsEmailValid(user.Email) {
		return "Invalid Email", false
	} else if len(strings.TrimSpace(user.Password)) == 0 {
		return "Password required", false
	} else if len(strings.TrimSpace(user.Password)) < 6 {
		return "Password should be at least 6 chars", false
	} else if len(strings.TrimSpace(user.PhoneNumber)) == 0 && flag != "login" {
		return "Phone num required", false
	} else if !utils.IsPhoneValid(user.PhoneNumber) && flag != "login" {
		return "Invalid Phonenumber", false
	}
	tempUser := &User{}
	if err := db.Raw("select * from users where email=?", user.Email).First(tempUser).Error; err != nil && err != gorm.ErrRecordNotFound {
		return "Internal server error", false
	}
	if flag == "login" {
		return "Requirement passed", true
	}
	if tempUser.Email != "" {
		return "User already exists", false
	}
	if err := db.Raw("select * from users where phoneNumber=?", user.PhoneNumber).First(tempUser).Error; err != nil && err != gorm.ErrRecordNotFound {
		return "Internal server error", false
	}
	if strings.TrimSpace(tempUser.PhoneNumber) == strings.TrimSpace(user.PhoneNumber) {
		return "Phone num already in use", false
	}
	return "Requirement passed", true
}
