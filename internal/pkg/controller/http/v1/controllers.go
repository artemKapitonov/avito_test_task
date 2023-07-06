package v1

import (
	"github.com/artemKapitonov/avito_test_task/internal/pkg/usecase"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	UseCase usecase.UseCase
}

func New(usecase *usecase.UseCase) *Controller {
	return &Controller{
		UseCase: *usecase,
	}
}

func (c *Controller) InitRoutes() *gin.Engine {
	router := gin.Default()

	account := router.Group("/account")
	{
		account.POST("/", c.createUser)
		account.GET("/:id", c.userByID)
		account.PUT("/:id", c.updateBalance)
	}
	return router

}
