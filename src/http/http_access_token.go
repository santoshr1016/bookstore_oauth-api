package http

import (
	"github.com/gin-gonic/gin"
	"github.com/santoshr1016/bookstore_oauth-api/src/domain/access_token"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func (hand *accessTokenHandler) GetById(ctx *gin.Context) {
	accessTokenId := strings.TrimSpace(ctx.Param("access_token_id"))
	accessToken, err := hand.service.GetById(accessTokenId)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusNotImplemented, accessToken)
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}
