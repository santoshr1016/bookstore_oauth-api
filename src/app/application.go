package app

import (
	"clients/cassandra"
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

	session, dbErr := cassandra.GetSession()
	if dbErr != nil {
		panic(dbErr)
	}
	session.Close()
	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token/", atHandler.Create)
	router.POST("/oauth/access_token/", atHandler.UpdateExpirationTime)
	router.Run(":8080")
}
