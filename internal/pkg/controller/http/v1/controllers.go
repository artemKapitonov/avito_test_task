package v1

import (
	"github.com/artemKapitonov/avito_test_task/internal/pkg/usecase"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Account
	Balance
	OperationHistory
}

func New(usecase *usecase.UseCase) *Controller {
	return &Controller{
		Account:          usecase.Account,
		Balance:          usecase.Balance,
		OperationHistory: usecase.OperationHistory,
	}
}

func (c *Controller) InitRoutes() *gin.Engine {
	router := gin.Default()

	account := router.Group("/account")
	{
		account.POST("/", c.createUser)
		account.GET("/:id", c.userByID)
		account.PUT("/:id", c.updateBalance)
		account.PUT("/transfer/:sender_id/:recipient_id", c.transfer)

	}
	return router

}
