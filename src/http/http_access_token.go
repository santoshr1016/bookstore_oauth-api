package http

import (
	"github.com/santoshr1016/bookstore_oauth-api/src/utils/errors"
	"github.com/santoshr1016/bookstore_oauth-api/src/domain/access_token"
	"net/http"
	"pkg/mod/github.com/gin-gonic/gin@v1.5.0"

	//"github.com/gin-gonic/gin"
	//"pkg/mod/github.com/gin-gonic/gin@v1.5.0"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(ctx *gin.Context)
	UpdateExpirationTime(ctx *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (hand *accessTokenHandler) GetById(ctx *gin.Context) {
	//accessTokenId := strings.TrimSpace(ctx.Param("access_token_id"))
	accessToken, err := hand.service.GetById(ctx.Param("access_token_id"))

	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusNotImplemented, accessToken)
}

func (hand *accessTokenHandler) Create(ctx *gin.Context) {
	var at access_token.AccessToken
	if err := ctx.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		ctx.JSON(restErr.Status, restErr)
		return
	}
	if err := hand.service.Create(at); err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusCreated, at)
}

func (hand *accessTokenHandler) UpdateExpirationTime(ctx *gin.Context) {
	ctx.JSON(http.StatusNotImplemented, "Not yet implemented")
}