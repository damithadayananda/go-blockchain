package response

type NodeResponse struct {
	SuccessResponse
	Result []string `json:"result"`
}
