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

type TinyGinServer struct {
	Port int `json:"port"`
}

func NewTinyGin(port int) TinyGinServer {
	tinyGin := TinyGinServer{Port: port}
	fmt.Println("\nStarting TinyGin...")
	return tinyGin
}

func (e TinyGinServer) Get(route string, handler Handler[any]) {
	http.HandleFunc(route, GetHttpHandler(handler))
	fmt.Println(fmt.Sprintf("[GET] %s", route))
}

func (e TinyGinServer) Run() {
	fmt.Println(fmt.Sprintf("\nServer running on: %d", e.Port))

	err := http.ListenAndServe(fmt.Sprintf(":%d", e.Port), nil)
	if err != nil {
		return
	}
}

func GetHttpHandler(handler Handler[any]) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		response := handler(Request{Query: r.URL.Query()})
		w.WriteHeader(response.HttpStatus)
		err := json.NewEncoder(w).Encode(response.Data)
		if err != nil {
			return
		}
	}
}
