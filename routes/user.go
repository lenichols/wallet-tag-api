package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lenichols/wallet-tag-api/controllers"
)

func UserRoute(router *gin.Engine) {
	// router.GET("/users/:id", func(c *gin.Context) {
	//     id := c.Param("id")
	//     c.JSON(200, gin.H{
	//         "id": id,
	//     })
	// })
	router.GET("/", controllers.GetUsersController)
	router.GET("/:id", controllers.GetUserController)
	router.POST("/", controllers.CreateUserController)
	router.DELETE("/:id", controllers.RemoveUserController)
	router.PUT("/:id", controllers.UpdateUserController)

}
