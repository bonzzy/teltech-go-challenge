package dtos

import "math/big"

type MathParams struct {
	X *big.Float `json:"x" form:"x" binding:"required,number"`
	Y *big.Float `json:"y" form:"y" binding:"required,number"`
}

type Response struct {
	Action string  `json:"action"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Answer float64 `json:"answer"`
	Cached bool    `json:"cached"`
}
