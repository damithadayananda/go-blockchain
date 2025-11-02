package domain

type Node struct {
	Ip          string
	Certificate []byte
	Address     string
	Status      NodeStatus
}

type NodeValidate struct {
	Address     string `json:"address"`
	Certificate []byte `json:"certificate"`
}
