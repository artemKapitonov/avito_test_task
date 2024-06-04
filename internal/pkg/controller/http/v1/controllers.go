package v1

import (
	"github.com/artemKapitonov/avito_test_task/internal/pkg/usecase"
	"github.com/artemKapitonov/avito_test_task/pkg/logging"
	"github.com/gin-gonic/gin"
)

// Controller represents the controller for handling HTTP requests.
type Controller struct {
	Account
	Balance
	OperationHistory
	CurrencyConverter
}

// New creates a new instance of the Controller.
func New(uc *usecase.UseCase) *Controller {
	return &Controller{
		Account:           uc.Account,
		Balance:           uc.Balance,
		OperationHistory:  uc.OperationHistory,
		CurrencyConverter: uc.CurrencyConverter,
	}
}

// InitRoutes initializes the routes for the controller.
func (c *Controller) InitRoutes(logger *logging.Logger) *gin.Engine {
	router := gin.New()

	router.Use(gin.LoggerWithWriter(logger.Writer))

	// Account routes
	account := c.RegisterAccountEndponts(router)

	c.RegisterHisotryEndpoints(account)

	return router
}

func (c *Controller) RegisterAccountEndponts(r *gin.Engine) *gin.RouterGroup {
	account := r.Group("/account")
	{
		account.PUT("/:id", c.updateBalance)
		account.PUT("/transfer/:sender_id/:recipient_id", c.transfer)
		account.POST("/", c.createUser)
		account.GET("/:id", c.userByID)
	}

	return account
}

func (c *Controller) RegisterHisotryEndpoints(account *gin.RouterGroup) *gin.RouterGroup {
	history := account.Group("/history")
	{
		history.GET("/:id", c.getHistory)
	}

	return history
}
