package routes

import (
	"bytes"
	"net/http"
	"testing"
)

func TestUrlRoutes(t *testing.T) {

	// Тестирование POST-запроса
	t.Run("PostUrlToGetShort", func(t *testing.T) {

		postData := []byte(`{"url": "http://example.com"}`)
		// Выполнение POST-запроса ко второму эндпоинту
		resp, err := http.Post("http://localhost:8080/", "application/json", bytes.NewBuffer(postData))
		if err != nil {
			t.Errorf(resp.Status)
			return
		}
		defer resp.Body.Close()

		// Здесь проверьте статус кода и тело ответа
		if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
			t.Errorf(resp.Status)
		}
	})

	// Тестирование GET-запроса
	t.Run("GetFullUrlByShort", func(t *testing.T) {
		// Выполнение GET-запроса к первому эндпоинту
		resp, err := http.Get("http://localhost:8080/qtj5opu")
		if err != nil {
			t.Errorf(resp.Status)
			return
		}
		defer resp.Body.Close()
	})
}