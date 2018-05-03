package watch

// EventType event type
type EventType string

const (
	// Added add
	Added EventType = "ADDED"
	// Modified modify
	Modified EventType = "MODIFIED"
	// Deleted delete
	Deleted EventType = "DELETED"
)
