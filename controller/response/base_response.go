package response

type BaseResponse struct {
	Success bool `json:"success"`
}
type SuccessResponse struct {
	BaseResponse
}

type FailResponse struct {
	BaseResponse
	Error string `json:"error"`
}
