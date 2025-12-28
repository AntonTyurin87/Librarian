package sources

import "time"

// Select ...
type Select struct {
	IDs       []int32
	Types     []string
	ObjectIDs []int32
	Address   []string
	CreatedAt []time.Time
}
