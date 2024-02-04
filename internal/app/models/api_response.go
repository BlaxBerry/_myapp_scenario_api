package models

type ApiResponseData struct {
	Code  int `json:"code"`
	Error any `json:"error"`
	Data  any `json:"data"`
}
