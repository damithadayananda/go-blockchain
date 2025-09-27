package domain

type NodeStatus int

const (
	ACTIVE NodeStatus = iota
	PENDING_VALIDATION
)

func (m NodeStatus) String() string {
	switch m {
	case ACTIVE:
		return "ACTIVE"
	case PENDING_VALIDATION:
		return "PENDING_VALIDATION"
	default:
		return "UNKNOWN"
	}
}
