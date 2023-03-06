package schema

type CycleState int

const (
	PREPARATION CycleState = iota + 1
	ACTIVE
	CLOSED
)
