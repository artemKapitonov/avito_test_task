package v1

import "github.com/artemKapitonov/avito_test_task/internal/pkg/usecase"

type Controller struct {
	UseCase usecase.UseCase
}

func New(usecase *usecase.UseCase) *Controller {
	return &Controller{
		UseCase: *usecase,
	}
}
