package v1

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/artemKapitonov/avito_test_task/internal/pkg/entity"
	"github.com/gin-gonic/gin"
)

type Account interface {
	Create(ctx context.Context) (entity.User, error)
	GetByID(ctx context.Context, id uint64) (entity.User, error)
}

func (c *Controller) createUser(ctx *gin.Context) {
	user, err := c.UseCase.Account.Create(ctx)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, fmt.Sprintf("Can't create user: %s", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *Controller) userByID(ctx *gin.Context) {
	param := ctx.Param("id")

	userID, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		errorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	user, err := c.UseCase.Account.GetByID(ctx, userID)
	if err != nil {
		errorResponse(ctx, http.StatusInternalServerError, fmt.Sprintf("Can't get user by id: %s", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
