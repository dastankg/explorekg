package app

import "github.com/gin-gonic/gin"

func App() *gin.Engine {
	router := gin.Default()

	return router
}
