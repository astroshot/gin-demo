package entity

import (
	"gin-demo/pkg/util"
)

// JSONResponse defines api return value
type JSONResponse struct {
	Code *int        `json:"code"`
	Info *string     `json:"info"`
	Data interface{} `json:"data"`
}

func Success(code int, info string, data interface{}) *JSONResponse {
	return &JSONResponse{
		Code: &code,
		Info: &info,
		Data: data,
	}
}

func SuccessWithDefaultCode(info string, data interface{}) *JSONResponse {
	return Success(util.DefaultSuccessCode, info, data)
}

func Fail(code int, info string, data interface{}) *JSONResponse {
	return &JSONResponse{
		Code: &code,
		Info: &info,
		Data: data,
	}
}

func FailWithDefaultCode(info string, data interface{}) *JSONResponse {
	return Fail(util.DefaultErrorCode, info, data)
}
