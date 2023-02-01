package core

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	Query map[string][]string
}

type Response struct {
	Data       any
	HttpStatus int
}

type Handler[T any] func(request Request) Response

type TinyGinRoute struct {
	HttpMethod string
	Route      string
	Handler    Handler[any]
}

type TinyGinServer struct {
	Port   int `json:"port"`
	Routes []TinyGinRoute
}

func NewTinyGin(port int) TinyGinServer {
	tinyGin := TinyGinServer{Port: port}
	return tinyGin
}

func (e *TinyGinServer) Get(route string, handler Handler[any]) {
	e.Routes = append(e.Routes, TinyGinRoute{Route: route, Handler: handler, HttpMethod: http.MethodGet})
}

func (e *TinyGinServer) Post(route string, handler Handler[any]) {
	e.Routes = append(e.Routes, TinyGinRoute{Route: route, Handler: handler, HttpMethod: http.MethodPost})
}

func (e *TinyGinServer) Run() {
	fmt.Println("\nStarting TinyGin...")
	fmt.Println(fmt.Sprintf("Server running on: %d", e.Port))
	fmt.Println(fmt.Sprintf("\nInitialiasing routes:"))

	routesByPath := groupRoutesByPath(e.Routes)

	for route, routeConfigList := range routesByPath {
		http.HandleFunc(route, GetHttpHandler(routeConfigList))
		for _, routeConfig := range routeConfigList {
			fmt.Println(fmt.Sprintf("[%s] %s", routeConfig.HttpMethod, routeConfig.Route))
		}
	}

	err := http.ListenAndServe(fmt.Sprintf(":%d", e.Port), nil)
	if err != nil {
		panic(err)
	}
}

// GetHttpHandler handling multiple Route paths with the same http method
func GetHttpHandler(routeConfigList []TinyGinRoute) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, routeConfig := range routeConfigList {
			if routeConfig.HttpMethod != r.Method {
				w.WriteHeader(http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			response := routeConfig.Handler(Request{Query: r.URL.Query()})
			w.WriteHeader(response.HttpStatus)
			err := json.NewEncoder(w).Encode(response.Data)

			if err != nil {
				return
			}
		}
	}
}

func groupRoutesByPath(routes []TinyGinRoute) map[string][]TinyGinRoute {
	routesByPath := make(map[string][]TinyGinRoute)

	for i := 0; i < len(routes); i++ {
		routeConfig := routes[i]
		if routesByPath[routeConfig.Route] != nil {
			routesByPath[routeConfig.Route] = append(routesByPath[routeConfig.Route], routeConfig)
			continue
		}
		routesByPath[routeConfig.Route] = []TinyGinRoute{routeConfig}
	}

	return routesByPath
}
