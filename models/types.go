package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phoneNumber"`
}

type Token struct {
	UserId uint   `json:"userId"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

type Menu struct {
	gorm.Model
	Category  string     `json:"category"`
	MenuItems []MenuItem `json:"menuItems" `
}

type MenuItem struct {
	gorm.Model
	MenuID   uint   `json:"menuId"`
	Name     string `json:"name"`
	Price    uint   `json:"price"`
	ImageUrl string `json:"imageUrl"`
}

type Table struct {
	gorm.Model
	TableNumber    uint    `json:"tableNumber"`
	Orders         []Order `json:"orders"`
	NumberOfGuests uint    `json:"numberOfGuests"`
}

type Order struct {
	gorm.Model
	TableID        uint       `json:"tableId"`
	OrderItems     []MenuItem `json:"orderItems"`
	OrderDate      time.Time  `json:"orderDate"`
	DeliveryStatus bool       `json:"deliveryStatus"`
}
