package handlers

type ApiResponseType struct {
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
	Code  int         `json:"code"`
}
