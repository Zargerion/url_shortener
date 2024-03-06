package routes

import (
	"github.com/Zargerion/url_shortener/internal/controller"
	"github.com/gin-gonic/gin"
)

func Url(router *gin.Engine, controller controller.UrlController) *gin.Engine {

	router.GET("/:url", controller.GetFullUrlByShort)
	router.POST("/", controller.PostUrlToGetShort)

	return router
}