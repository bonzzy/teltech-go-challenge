package validations

import "github.com/bonzzy/teltech-go-challenge/core"

func ValidateQueryXY(query map[string][]string) (bool, string) {
	if len(query["x"]) == 0 || !core.IsNumber(query["x"][0]) {
		return false, "x needs to be a number!"
	}
	if len(query["y"]) == 0 || !core.IsNumber(query["y"][0]) {
		return false, "y needs to be a number!"
	}

	return true, ""
}
