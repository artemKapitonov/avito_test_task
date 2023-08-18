package v1

import (
	"github.com/artemKapitonov/avito_test_task/internal/pkg/usecase"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Account
	Balance
	OperationHistory
	CurrencyConverter
}

func New(uc *usecase.UseCase) *Controller {
	return &Controller{
		Account:           uc.Account,
		Balance:           uc.Balance,
		OperationHistory:  uc.OperationHistory,
		CurrencyConverter: uc.CurrencyConverter,
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

		history := account.Group("/history")
		{
			history.GET("/:id", c.getHistory)
		}
	}
	return router

}
