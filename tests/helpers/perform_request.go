package helpers

import (
	"github.com/bonzzy/teltech-go-challenge/core"
	"net/http"
	"net/http/httptest"
)

func PerformRequest(handler core.Handler[any], method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()

	routeConfig := core.TinyGinRoute{
		Route:      path,
		HttpMethod: method,
		Handler:    handler,
	}
	core.GetHttpHandler([]core.TinyGinRoute{routeConfig})(w, req)
	return w
}
