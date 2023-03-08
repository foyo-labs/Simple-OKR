package schema

type Objective struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	UserID      string          `json:"user_id"`
	Description string          `json:"discription"`
	Actived     ObjectiveStatus `json:"actived"`
	Sequence    int32           `json:"sequence"`
	ParentID    int32           `json:"parent_id"`
	Created     uint64          `json:"created"`
	KeyResults  []*KeyResult    `json:"key_results"`
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
	Name        string             `json:"name"`
	Description string             `json:"discription"`
	ParentID    int32              `json:"parent_id"`
	KeyResults  []*ReqestKeyResult `json:"key_results"`
}

type ReqestKeyResult struct {
	Name string `json:"name"`
}
