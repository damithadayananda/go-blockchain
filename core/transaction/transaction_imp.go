package transaction

import "go-blockchain/domain"

type TransactionInterface interface {
	SetMiningStatus(state domain.MiningStates)
}
