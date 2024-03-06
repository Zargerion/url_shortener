package layers

import "github.com/Zargerion/url_shortener/internal/controller"

type IController struct {
	UrlController controller.UrlController
}

func Controller(m *IModel) *IController {
	return &IController{
		UrlController: controller.NewUrlController(m.UrlModel),
	}
}
