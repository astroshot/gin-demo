package model

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

func Fail(code int, info string, data interface{}) *JSONResponse {
	return &JSONResponse{
		Code: &code,
		Info: &info,
		Data: data,
	}
}
