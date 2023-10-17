package controllers

import (
	"apna-restaurant/database"
	"apna-restaurant/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	database.DBConfig()
	db = database.GetDB()
}

func AddMenu(c *gin.Context) {
	menu := &models.Menu{}
	if err := c.ShouldBindJSON(menu); err != nil {
		fmt.Println("returning from here", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	// db := database.GetDB()
	if resp, ok := menu.ValidateMenuAddRequest(db, ""); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp})
		return
	}
	tempMenu := &models.Menu{}
	err := db.Raw("select * from menus where id=?", menu.ID).First(tempMenu).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if tempMenu.ID > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Menu id already exists"})
		return
	}
	db.Create(menu)
	if menu.ID <= 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"meesage": "Menu added"})
}

func GetAllMenus(c *gin.Context) {
	fmt.Println("came inside")
	menus := &[]models.Menu{}
	err := db.Raw("select * from menus").Find(&menus).Error
	if err != nil {
		fmt.Println("came here", err)
		c.JSON(http.StatusOK, gin.H{"error": "Internal server error"})
		return
	}
	if len(*menus) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No menus found, please add one"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": menus})
}

func GetMenuByID(c *gin.Context) {
	id, isIdInput := c.GetQuery("id")
	if !isIdInput {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing menu id"})
		return
	}
	// db := database.GetDB()
	resultantMenu := &models.Menu{}
	err := db.Raw("select * from menus where id=?", id).First(resultantMenu).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": resultantMenu})
}

func UpdateMenu(c *gin.Context) {

	reqMenu := &models.Menu{}
	if err := c.ShouldBindJSON(reqMenu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	// db := database.GetDB()

	if resp, ok := reqMenu.ValidateMenuAddRequest(db, "update"); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp})
		return
	}
	resultantMenu := &models.Menu{}
	err := db.Raw("select * from menus where id=?", reqMenu.ID).First(resultantMenu).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	resultantMenu.Category = reqMenu.Category
	resultantMenuItems := append(resultantMenu.MenuItems, reqMenu.MenuItems...)
	resultantMenu.MenuItems = resultantMenuItems
	db.Save(resultantMenu)
	c.JSON(http.StatusOK, gin.H{"updated_data": resultantMenu})
}
