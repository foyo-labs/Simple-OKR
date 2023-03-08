package schema

type CycleState int

const (
	PREPARATION CycleState = iota + 1
	ACTIVE
	CLOSED
)

type Role int

const (
	AdminUser Role = iota
	LocalUser
)

type ObjectiveStatus int

const (
	ObjectiveActived ObjectiveStatus = iota + 1
	ObjectiveDisabled
)

type ObjectiveType int

const (
	GroupObjectiveType ObjectiveType = iota + 1
	UserObjectiveType
)
