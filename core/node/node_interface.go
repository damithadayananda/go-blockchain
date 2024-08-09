package node

type NodeInterface interface {
	GetNodes() ([]string, error)
	SaveNode(url string) error
	RemoveNode(url string) error
}
