package controllers

import (
	"apna-restaurant/database"
	"apna-restaurant/middleware"
	"apna-restaurant/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Signup(c *gin.Context) {
	user := &models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}
	db := database.GetDB()
	if resp, ok := user.ValidateUserAddRequest(db, ""); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error hashing password"})
		return
	}
	user.Password = string(hashedPassword)
	db.Create(user)
	if user.ID <= 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"insertion_id": user.ID})
}

func Signin(c *gin.Context) {
	user := &models.User{}
	response := make(map[string]interface{})
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}
	db := database.GetDB()
	if resp, ok := user.ValidateUserAddRequest(db, "login"); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": resp})
		return
	}
	tempUser := &models.User{}
	if err := db.Raw("select * from users where email=?", user.Email).First(tempUser).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"message": "User does not exist"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal server error"})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(tempUser.Password), []byte(user.Password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid credentials"})
			return
		}
	}
	token := middleware.GenerateToken(tempUser.ID, tempUser.Email)
	data := map[string]interface{}{
		"user":  tempUser,
		"token": token,
	}
	response["data"] = data
	c.JSON(http.StatusCreated, response)
}
