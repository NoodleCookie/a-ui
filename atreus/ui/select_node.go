package ui

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

type SelectNode struct {
	id       string
	ui       *promptui.Select
	bindData *[]string
	callback func(key string) error
	errCatch func(err error)
}

func (sn *SelectNode) Process() chan process {
	return nil
}

func (sn *SelectNode) Registry(ui *UI) {
	ui.RegistryNode(sn)
}

func (sn *SelectNode) ID(id string) *SelectNode {
	sn.id = id
	return sn
}

func (sn *SelectNode) Label(label string) *SelectNode {
	if sn.ui == nil {
		sn.ui = &promptui.Select{}
	}
	sn.ui.Label = label
	return sn
}

func (sn *SelectNode) Bind(data *[]string) *SelectNode {
	if sn.ui == nil {
		sn.ui = &promptui.Select{}
	}
	sn.bindData = data
	sn.ui.Items = *data
	return sn
}

func (sn *SelectNode) Then(callback func(key string) error) *SelectNode {
	sn.callback = callback
	return sn
}

func (sn *SelectNode) Catch(errCatch func(err error)) *SelectNode {
	sn.errCatch = errCatch
	return sn
}

func (sn *SelectNode) GetID() string {
	return sn.id
}

func (sn *SelectNode) SecurityCheck() error {
	if sn.ui == nil || sn.ui.Items == nil {
		return fmt.Errorf("[Select_Node_%s] require ui and ui items", sn.id)
	}
	if sn.errCatch == nil {
		return fmt.Errorf("[Select_Node_%s] require error function", sn.id)
	}
	if sn.callback == nil {
		return fmt.Errorf("[Select_Node_%s] require callback function", sn.id)
	}
	return nil
}

func (sn *SelectNode) Do() {
	sn.ui.Items = *sn.bindData
	_, key, err := sn.ui.Run()
	if err != nil {
		sn.errCatch(err)
	}
	if err := sn.callback(key); err != nil {
		sn.errCatch(err)
	}
}
