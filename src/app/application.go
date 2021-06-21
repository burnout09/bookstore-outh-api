package app

import (
	"github.com/burnout09/bookstore-outh-api/src/domain/access_token"
	"github.com/burnout09/bookstore-outh-api/src/http"
	"github.com/burnout09/bookstore-outh-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var(
	router=gin.Default()
)

func StartApplication() {
	//session, dbErr := cassandra.GetSession()
	//if dbErr != nil {
	//	panic(dbErr)
	//}
	//session.Close()

	atHandler:=http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	router.Run(":8080")
}
