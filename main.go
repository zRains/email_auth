package main

import (
	"context"
	"email_auth/data"
	"email_auth/initer"
	"email_auth/middle"
	"email_auth/router"
	"email_auth/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	ctx := context.TODO()

	// Create gin server
	server := gin.Default()

	// Gin cors middlewares
	server.Use(middle.Cors())
	//server.Use(middle.TokenAuth())

	// New api router group
	api := server.Group("/api")

	// Init
	initer.InitConfig(".")
	initer.InitDatabase(ctx)
	data.InitUserCollection()
	router.InitCommonRouter(api)
	router.InitUserRouter(ctx, service.CreateUserService(ctx), api)
	defer initer.CloseDatabase(ctx)

	log.Fatal(server.Run())
}
