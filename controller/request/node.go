package request

type AddNodeRequest struct {
	Url           string   `json:"url"`
	InformedNodes []string `json:"informed_nodes"`
}
