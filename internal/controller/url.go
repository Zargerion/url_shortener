package controller

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Zargerion/url_shortener/internal/model"
)

type UrlController struct {
	m model.UrlModel
}

func NewUrlController(m model.UrlModel) UrlController {
	return UrlController{m: m}
}

func (uc *UrlController) GetFullUrlByShort(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	url, ok := c.Params.Get("url")
	if !ok || url == "" {
		c.String(http.StatusBadRequest, "Не взялись параметры.")
		return
	}

	ans, err := uc.m.GetFullUrlByShort(ctx, url)
	if err != nil {
		c.String(http.StatusBadRequest, "Не выполнилась модель действий.")
		return
	}

	c.String(http.StatusOK, ans)
}


func (uc *UrlController) PostUrlToGetShort(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var url string
	if err := c.ShouldBind(&url); err != nil {
		c.String(http.StatusBadRequest, "Не удалось прочитать тело запроса.")
		return
	}

	ans, err := uc.m.PostUrlToGetShort(ctx, url)
	if err != nil {
		c.String(http.StatusBadRequest, "Не выполнилась модель действий.")
		return
	}

	c.String(http.StatusOK, ans)
}
