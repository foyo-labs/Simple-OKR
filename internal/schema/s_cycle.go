package schema

type Cycle struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	StartAt uint64 `json:"start_at"`
	EndAt   uint64 `json:"end_at"`
}

type CreateCycleRequest struct {
	Name    string `json:"name"`
	StartAt uint64 `json:"start_at"`
	EndAt   uint64 `json:"end_at"`
}
