package app

import (
	"explorekg/config"
	"explorekg/internal/auth"
	"explorekg/internal/user"
	"explorekg/pkg/db"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func App() *gin.Engine {
	router := gin.Default()
	conf := config.LoadConfig()
	db, err := db.NewDb(conf)
	if err != nil {
		log.Fatal(err)
	}
	userRepository := user.NewUserRepository(db)
	router.Use(gin.Logger())
	url := ginSwagger.URL("/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))
	authService := auth.NewAuthService(userRepository)
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Config:      conf,
		AuthService: authService,
	})
	return router
}
