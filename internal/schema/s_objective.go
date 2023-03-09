package schema

type Objective struct {
	ID            string          `json:"id"`
	Name          string          `json:"name"`
	UserID        string          `json:"user_id"`
	Description   string          `json:"discription"`
	Actived       ObjectiveStatus `json:"actived"`
	Sequence      int32           `json:"sequence"`
	ParentID      int32           `json:"parent_id"`
	Created       uint64          `json:"created"`
	Updated       uint64          `json:"updated"`
	KeyResults    []*KeyResult    `json:"key_results"`
	ObjectiveType ObjectiveType   `json:"objective_type"`
	GroupID       string          `json:"group_id"`
	CycleID       string          `json:"cycle_id"`
}

type KeyResult struct {
	ID           string  `json:"id"`
	ObjectiveID  string  `json:"objective_id"`
	Name         string  `json:"name"`
	StartValue   float64 `json:"start_value"`
	TargetValue  float64 `json:"target_value"`
	CurrentValue float64 `json:"current_value"`
	Sequence     int32   `json:"sequence"`
	Created      uint64  `json:"created"`
}

type KeyResults []*KeyResult

type ReqestObjective struct {
	Name          string             `json:"name"`
	Description   string             `json:"discription"`
	ParentID      int32              `json:"parent_id"`
	KeyResults    []*ReqestKeyResult `json:"key_results"`
	ObjectiveType ObjectiveType      `json:"objective_type"`
	GroupID       string             `json:"group_id"`
	CycleID       string             `json:"cycle_id"`
}

type ReqestKeyResult struct {
	Name string `json:"name"`
}

type UserObjective struct {
	ID          string `json:"id"`
	UserID      string `json:"user_id"`
	ObjectiveID string `json:"objective_id"`
	CycleID     string `json:"cycle_id"`
}

type GroupObjective struct {
	ID          string `json:"id"`
	GroupID     string `json:"group_id"`
	ObjectiveID string `json:"objective_id"`
	CycleID     string `json:"cycle_id"`
}

type ObjectiveQueryParam struct {
	CycleID string `json:"cycle_id"`
	GroupID string `json:"group_id"`
}
