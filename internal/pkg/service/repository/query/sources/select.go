package sources

import "time"

// Select ...
type Select struct {
	IDs         []int32
	Types       []int32
	ObjectIDs   []int32
	Address     []string
	CreatedAt   []time.Time
	TimePeriods []string
}
