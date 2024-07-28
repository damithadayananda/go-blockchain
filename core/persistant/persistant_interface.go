package persistant

import "go-blockchain/core"

type Persister interface {
	Save(value core.Block) error
	GetAll() ([]core.Block, error)
	GetLastBlock() (core.Block, error)
}
