package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lenichols/wallet-tag-api/config"
	"github.com/lenichols/wallet-tag-api/models"
)

func GetUsersController(c *gin.Context) {
	// c.String(200, "User Authenticated")
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)

}

func GetUserController(c *gin.Context) {
	// c.String(200, "User Authenticated")
	users := []models.User{}
	config.DB.Where("ID = ?", c.Param("id")).Find(&users)
	c.JSON(200, &users)

}

func CreateUserController(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, &user)
}

func RemoveUserController(c *gin.Context) {
	var user models.User
	config.DB.Where("ID = ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)
}

func UpdateUserController(c *gin.Context) {
	var user models.User
	config.DB.Where("ID = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}
