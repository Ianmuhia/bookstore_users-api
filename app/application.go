package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()

	)
func StartApplication() {
	mapUrls()
	err := router.Run("127.0.0.1:8090")
	if err != nil {
		return 
	}
}