package server

import (
	"os"

	"github.com/gin-gonic/gin"
)

func NewGin() *gin.Engine {
	app := gin.Default()
	if os.Getenv("GIN_MODE") == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Установка максимального размера тела запроса в 40MB
	app.MaxMultipartMemory = 40 * 1024 * 1024

	return app
}
