package core

import "strconv"

func IsNumber(value string) bool {
	_, err := strconv.ParseFloat(value, 64)
	return err == nil
}
