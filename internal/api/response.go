package api

import "encoding/json"

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
}

func NewErrorResponse(msg string) []byte {
	resp := ErrorResponse{Error: msg}
	bytez, _ := json.Marshal(resp)

	return bytez
}

func NewSuccessResponse(data interface{}) []byte {
	resp := SuccessResponse{Data: data}
	bytez, _ := json.Marshal(resp)

	return bytez
}
