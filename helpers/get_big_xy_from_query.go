package helpers

import (
	"math/big"
)

func GetBigXYFromQuery(query map[string][]string) (*big.Float, *big.Float) {
	x, _, _ := new(big.Float).SetPrec(256).SetMode(big.ToZero).Parse(query["x"][0], 10)
	y, _, _ := new(big.Float).SetPrec(256).SetMode(big.ToZero).Parse(query["y"][0], 10)

	return x, y
}
