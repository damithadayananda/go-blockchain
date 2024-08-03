package domain

type MiningStates int

const (
	READY_FOR_MINING MiningStates = iota
	MINING_DONE
	READY_FOR_VALIDATION
)
