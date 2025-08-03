package inmemorydb

import (
	"go-blockchain/domain"
)

type InMemoryNode struct {
	nodes []domain.Node
}

func NewInMemoryNode() *InMemoryNode {
	return &InMemoryNode{}
}

func (n *InMemoryNode) Save(node domain.Node) error {
	n.nodes = append(n.nodes, node)
	return nil
}
func (n *InMemoryNode) Delete(ip string) error {
	for k, v := range n.nodes {
		if v.Ip == ip {
			n.nodes = append(n.nodes[:k], n.nodes[k+1:]...)
		}
	}
	return nil
}
func (n *InMemoryNode) GetAll() ([]domain.Node, error) {
	return n.nodes, nil
}
