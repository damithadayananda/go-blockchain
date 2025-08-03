package request

import "go-blockchain/domain"

type AddNodeRequest struct {
	Url           string        `json:"url"`
	InformedNodes []domain.Node `json:"informed_nodes"`
	Certificate   []byte        `json:"certificate"`
	Address       string        `json:"address"`
}
