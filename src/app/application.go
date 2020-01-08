package app

import (
	"github.com/gin-gonic/gin"
	"github.com/santoshr1016/bookstore_oauth-api/src/domain/access_token"
	"github.com/santoshr1016/bookstore_oauth-api/src/http"
	"github.com/santoshr1016/bookstore_oauth-api/src/repository/db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	//dbRepo := db.NewRepository()
	//// Access token service needs the db repo
	//atService := access_token.NewService(dbRepo)
	//atHandler := http.NewHandler(atService)

	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.Run(":8080")
}
