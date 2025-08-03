package response

type NodeResponse struct {
	SuccessResponse
	Result []Node `json:"result"`
}

type Node struct {
	Ip          string
	Certificate []byte
	Address     string
}
