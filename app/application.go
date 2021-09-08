package app

import (
	"bookstore_users-api/logger"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Log.Info("about to start application")
	err := router.Run("127.0.0.1:8090")
	if err != nil {
		return
	}
}
