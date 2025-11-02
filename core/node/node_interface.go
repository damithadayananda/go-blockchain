package node

import (
	"go-blockchain/core/transaction"
	"go-blockchain/domain"
)

type NodeInterface interface {
	GetNodes() ([]domain.Node, error)
	SaveNode(node domain.Node) error
	RemoveNode(url string) error
	ValidateNewNode(validate domain.NodeValidate) (transaction.Transaction, error)
}
