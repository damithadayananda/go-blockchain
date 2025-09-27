package domain

type MiningStates int

const (
	READY_FOR_MINING MiningStates = iota
	MINING_DONE
	READY_FOR_VALIDATION
)

func (m MiningStates) String() string {
	switch m {
	case READY_FOR_MINING:
		return "READY_FOR_MINING"
	case MINING_DONE:
		return "MINING_DONE"
	case READY_FOR_VALIDATION:
		return "READY_FOR_VALIDATION"
	default:
		return "UNKNOWN"
	}
}
