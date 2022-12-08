package ui

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

type InputNode struct {
	id       string
	ui       *promptui.Prompt
	callback func(input string) error
	errCatch func(err error)
}

func (e *InputNode) Process() chan process {
	return nil
}

func (e *InputNode) Registry(ui *UI) {
	ui.RegistryNode(e)
}

func (e *InputNode) ID(id string) *InputNode {
	e.id = id
	return e
}

func (e *InputNode) Label(label string) *InputNode {
	if e.ui == nil {
		e.ui = &promptui.Prompt{}
	}
	e.ui.Label = label
	return e
}

func (e *InputNode) GetID() string {
	return e.id
}

func (e *InputNode) Do() {
	if e.ui == nil {
		e.ui = &promptui.Prompt{}
	}
	input, err := e.ui.Run()
	if err != nil {
		e.errCatch(err)
	}
	if err := e.callback(input); err != nil {
		e.errCatch(err)
	}
}

func (e *InputNode) SecurityCheck() error {
	if e.id == "" {
		return fmt.Errorf("[INPUT_NODE] require id with .ID()")
	}
	if e.callback == nil {
		return fmt.Errorf("[INPUT_NODE_%s] require callback method with .Then()", e.id)
	}
	if e.errCatch == nil {
		return fmt.Errorf("[INPUT_NODE_%s] require catch error method with .Catch()", e.id)
	}
	return nil
}

func (e *InputNode) Then(callback func(input string) error) *InputNode {
	e.callback = callback
	return e
}

func (e *InputNode) Catch(errCatch func(err error)) *InputNode {
	e.errCatch = errCatch
	return e
}
