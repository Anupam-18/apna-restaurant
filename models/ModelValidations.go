package models

import (
	"apna-restaurant/utils"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func (user *User) ValidateUserAddRequest(db *gorm.DB, flag string) (string, bool) {
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
		fmt.Println(err)
		return "Internal server error", false
	}
	if flag == "login" {
		return "Requirement passed", true
	}
	if tempUser.Email != "" {
		return "User already exists", false
	}
	if err := db.Raw("select * from users where phone_number=?", user.PhoneNumber).First(tempUser).Error; err != nil && err != gorm.ErrRecordNotFound {
		return "Internal server error", false
	}
	if strings.TrimSpace(tempUser.PhoneNumber) == strings.TrimSpace(user.PhoneNumber) {
		return "Phone num already in use", false
	}
	return "Requirement passed", true
}

func (item *MenuItem) ValidateMenuItem() bool {
	if len(strings.TrimSpace(item.Name)) < 3 {
		return false
	} else if item.Price < 0 {
		return false
	} else if !utils.IsURLValid(item.ImageUrl) {
		return false
	}
	return true
}

func (menu *Menu) ValidateMenuAddRequest(db *gorm.DB, flag string) (string, bool) {
	if flag != "" {
		if menu.ID < 0 {
			return "Menu Id required", false
		}
		tempMenu := &Menu{}
		err := db.Raw("select * from menus where id=?", menu.ID).First(tempMenu).Error
		if err != nil {
			return "Internal server error", false
		}
		if tempMenu.ID < 0 {
			return "Menu does not exist", false
		}
	} else if len(strings.TrimSpace(menu.Category)) == 0 {
		return "Category required", false
	} else if len(strings.TrimSpace(menu.Category)) < 4 {
		return "Category should be at least 4 chars longer", false
	} else if len(menu.MenuItems) == 0 {
		return "Menu cannot be empty, Please add an item", false
	}
	for _, item := range menu.MenuItems {
		if !item.ValidateMenuItem() {
			return "One of menutem is invalid", false
		}
	}
	return "Requirement passed", true
}
