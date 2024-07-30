package response

type SuccessResponse struct {
	Success bool `json:"success"`
}

type FailResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
