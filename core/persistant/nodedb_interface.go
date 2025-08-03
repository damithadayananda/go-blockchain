package persistant

import (
	"go-blockchain/domain"
)

type NodeDBInterface interface {
	Save(node domain.Node) error
	GetAll() ([]domain.Node, error)
	Delete(ip string) error
}
