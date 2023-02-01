package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/bonzzy/teltech-go-challenge/core"
	"github.com/bonzzy/teltech-go-challenge/dtos"
	"github.com/bonzzy/teltech-go-challenge/services"
	"net/http"
	"sort"
)

func CacheWrapper(controllerHandler core.Handler[any], responseType dtos.Response) func(request core.Request) core.Response {
	return func(request core.Request) core.Response {
		cacheKey := getCacheKeyFromQuery(request.Query)

		fmt.Println(fmt.Sprintf("Cache key: %s", cacheKey))
		cacheValue, err := services.GetCacheValue(cacheKey)

		if err == nil {
			fmt.Println(fmt.Sprintf("Cache Hit: %s", cacheValue))

			err := json.Unmarshal([]byte(cacheValue), &responseType)
			if err != nil {
				// panic
			}

			fmt.Println(fmt.Sprintf("Cache Hit with unmarshal: %+v", responseType))

			responseType.Cached = true
			return core.Response{
				Data:       responseType,
				HttpStatus: http.StatusOK,
			}
		}

		fmt.Println("Cache miss")

		handlerResponse := controllerHandler(request)
		if handlerResponse.HttpStatus == http.StatusBadRequest {
			return handlerResponse
		}

		handlerResponseJson, _ := json.Marshal(handlerResponse.Data)

		services.SetCacheValue(cacheKey, string(handlerResponseJson))
		return handlerResponse
	}
}

func getCacheKeyFromQuery(query map[string][]string) string {
	var cacheKey string

	// sort keys in the query because
	// range query doesn't run in the same order every time
	params := make([]string, 0, len(query))
	for k := range query {
		params = append(params, k)
	}
	sort.Strings(params)

	for _, param := range params {
		cacheKey += param
		paramValues := query[param]
		for i := 0; i < len(paramValues); i++ {
			cacheKey += paramValues[i]
		}
	}

	return cacheKey
}
