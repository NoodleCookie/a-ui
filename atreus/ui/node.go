package ui

type Node interface {
	// GetID get node unique key
	GetID() string

	// Do main logic
	Do()

	// SecurityCheck check node's format
	SecurityCheck() error

	// Registry this node to an ui framework
	Registry(ui *UI)

	// Process return this node's process channel
	Process() chan process
}

var _ Node = (*SelectNode)(nil)
var _ Node = (*EmptyNode)(nil)
var _ Node = (*InputNode)(nil)
var _ Node = (*BusinessSelectNode)(nil)
