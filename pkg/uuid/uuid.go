package uuid

import (
	"fmt"

	"github.com/google/uuid"
)

// UUID Define alias
type UUID = uuid.UUID

// NewUUID Create uuid
func NewUUID() (UUID, error) {
	return uuid.NewRandom()
}

// MustUUID Create uuid(Throw panic if something goes wrong)
func MustUUID() UUID {
	v, err := NewUUID()
	if err != nil {
		panic(err)
	}
	return v
}

// MustString Create uuid
func MustString() string {
	return MustUUID().String()
}

func NextID() string {
	currWoker := &IdWorker{}
	currWoker.InitIdWorker(1000, 1)
	newId, _ := currWoker.NextId()
	return fmt.Sprintf("%d", newId)
}
