package handlers

type ApiResponseType struct {
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
	Code  int         `json:"code"`
}
