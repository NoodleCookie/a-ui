package ui

import (
	"fmt"

	"github.com/pkg/errors"
)

type UI struct {
	nodes   map[string]Node
	process chan process
	storage map[string]interface{}
}

// RegistryNode registry node into ui framework with its id
func (ui *UI) RegistryNode(node Node) *UI {
	if ui.nodes == nil {
		ui.nodes = make(map[string]Node, 0)
	}
	if _, ok := ui.nodes[node.GetID()]; ok {
		fmt.Printf("node_id %s have been registried\n", node.GetID())
	}
	ui.nodes[node.GetID()] = node
	return ui
}

// GetNodeByID get registry node from ui framework, otherwise return error
func (ui *UI) GetNodeByID(id string) (node Node, err error) {
	if node, ok := ui.nodes[id]; ok {
		return node, nil
	} else {
		return nil, fmt.Errorf("node_id %s does not exist", id)
	}
}

// SetLocalStorage set a value in ui storage with key
func (ui *UI) SetLocalStorage(key string, value interface{}) {
	if ui.storage == nil {
		ui.storage = make(map[string]interface{}, 0)
	}
	ui.storage[key] = value
}

// LocalStorage return a value defined in ui storage by key
func (ui *UI) LocalStorage(key string) (interface{}, error) {
	if ui.storage == nil {
		return nil, fmt.Errorf("nothing found in ui storage")
	}
	value, ok := ui.storage[key]
	if !ok {
		return nil, fmt.Errorf("nothing found in ui storage")
	}
	return value, nil
}

// Go specify a(some) defined node(s) to start ui if this node security check succeed
func (ui *UI) Go(ids ...string) (err error) {
	for _, id := range ids {
		if node, ok := ui.nodes[id]; ok {
			if err = node.SecurityCheck(); err != nil {
				return err
			}
			if node.Process() != nil {
				ui.process = node.Process()
			}
			node.Do()
			return nil
		} else {
			return fmt.Errorf("ui go node failed, node_id %s does not exist", id)
		}
	}
	return nil
}

// Process return ui current node's process channel, receive Awake to execute bind method, Stop to continue
func (ui *UI) Process() chan process {
	return ui.process
}

// Awake the ui block thread and waiting for block process to continue
func (ui *UI) Awake() {
	if ui.process != nil {
		ui.process <- Awake
		select {
		case <-ui.process:
		}
	}
}

func (ui *UI) securityCheck() (err error) {
	if ui.nodes == nil {
		return nil
	}
	for id, node := range ui.nodes {
		if err = node.SecurityCheck(); err != nil {
			return errors.Wrap(err, fmt.Sprintf("node_%s security check failed", id))
		}
	}
	return nil
}
