package ui

import (
	"fmt"
)

type EmptyNode struct {
	id       string
	callback func() error
	errCatch func(err error)
}

func (e *EmptyNode) Process() chan process {
	return nil
}

func (e *EmptyNode) Registry(ui *UI) {
	ui.RegistryNode(e)
}

func (e *EmptyNode) ID(id string) *EmptyNode {
	e.id = id
	return e
}

func (e *EmptyNode) Then(callback func() error) *EmptyNode {
	e.callback = callback
	return e
}

func (e *EmptyNode) Catch(errCatch func(err error)) *EmptyNode {
	e.errCatch = errCatch
	return e
}

func (e *EmptyNode) GetID() string {
	return e.id
}

func (e *EmptyNode) Do() {
	err := e.callback()
	if e.errCatch != nil {
		e.errCatch(err)
		return
	}
	panic(err)
}

func (e *EmptyNode) SecurityCheck() error {
	if e.callback == nil {
		return fmt.Errorf("[EMPTY_NODE_%s] need callback method", e.id)
	}
	return nil
}
