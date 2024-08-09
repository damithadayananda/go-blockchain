package node

import "go-blockchain/core/persistant"

var NodeRef *Node

type Node struct {
	database persistant.NodeDBInterface
}

func NewNode(database persistant.NodeDBInterface) {
	NodeRef = &Node{
		database: database,
	}
}

func (n *Node) SaveNode(url string) error {
	return n.database.Save(url)
}

func (n *Node) GetNodes() ([]string, error) {
	return n.database.GetAll()
}

func (n *Node) RemoveNode(url string) error {
	return n.database.Delete(url)
}
