package v1

import "github.com/gin-gonic/gin"

func (c *Controller) InitRoutes() *gin.Engine {
	router := gin.Default()

	account := router.Group("/account")
	{
		account.POST("/", c.createUser)
		account.GET("/:id", c.userByID)
	}
	return router

}
