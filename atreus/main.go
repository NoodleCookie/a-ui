package main

import (
	"fmt"
	"time"
	"ui/atreus/ui"
)

func main() {
	//SimpleDemo()
	//AsyncDemo()
	BusinessDemo()
}

func SimpleDemo() {

	// 1. create ui
	framework := new(ui.UI)

	// 2. define ui nodes id
	const (
		mainId            = "main"
		chooseGroupNodeId = "chooseGroupNode"
		inputBodyNodeId   = "inputBodyNodeId"
	)

	// 3. define ui node bind data
	var (
		testSuites []string
	)

	// 4. declaration nodes
	{
		node := ui.SelectNode{}
		nodeData := make([]string, 0, 3)
		const (
			selectGroup = "select group"
			freeMode    = "free mode"
			exit        = "exit"
		)
		nodeData = append(nodeData, selectGroup, freeMode, exit)
		node.ID(mainId).Label("select to continue").Bind(&nodeData).Then(func(key string) error {
			switch key {
			case selectGroup:
				fmt.Println("select group =>")
				testSuites = []string{"test_01", "test_02"}
				return framework.Go(chooseGroupNodeId)
			case freeMode:
				fmt.Println("free mode =>")
				return nil
			default:
				fmt.Println("exit =>")
				return nil
			}
		}).Catch(func(err error) { fmt.Println("catch error => ", err.Error()) }).Registry(framework)
	}

	{
		node := ui.SelectNode{}
		node.ID(chooseGroupNodeId).Label("choose your group to test").Bind(&testSuites).Then(func(key string) error {
			fmt.Println("execute test... => ", key)
			return framework.Go(inputBodyNodeId)
		}).Catch(func(err error) {
			fmt.Println("catch error => ", err.Error())
		}).Registry(framework)
	}

	{
		node := ui.InputNode{}
		node.ID(inputBodyNodeId).Label("input your body").Then(func(input string) error {
			fmt.Println("you have input value: ", input)
			return framework.Go(mainId)
		}).Catch(func(err error) {
			fmt.Println("catch error => ", err.Error())
		}).Registry(framework)
	}

	// 5. start your ui
	err := framework.Go(mainId)
	if err != nil {
		panic(err)
	}
}

func AsyncDemo() {
	framework := new(ui.UI)
	const (
		mainNodeId = "mainNodeId"
	)

	var (
		mainNodeData = []string{"test", "exit"}
	)

	{
		node := ui.SelectNode{}
		node.ID(mainNodeId).Label("== Async demo ==").Bind(&mainNodeData).Then(func(key string) error {
			switch key {
			case "test":
				go func() {
					time.Sleep(2 * time.Second)
					fmt.Println("go runtine awake")
				}()
				time.Sleep(10 * time.Second)
				fmt.Println("go main awake")
				return nil
			default:
				return nil
			}
		}).Catch(func(err error) {
			fmt.Println("error => ", err.Error())
		}).Registry(framework)
	}

	err := framework.Go(mainNodeId)
	if err != nil {
		panic(err)
	}
}

func BusinessDemo() {
	framework := new(ui.UI)

	const (
		mainId = "mainId"
	)

	var (
		mainNodeData = []string{"1", "2"}
	)

	{
		node := ui.BusinessSelectNode{}
		node.
			Mode(ui.Mix).
			Business(func() error {
				fmt.Printf("Business => %s\n", "so bussy")
				return nil
			}).
			Bind(&mainNodeData).
			Then(func(key string) error {
				framework.Awake()
				fmt.Printf("select => %s\n", key)
				return nil
			}).
			Catch(func(err error) {
				fmt.Printf("error => %s\n", err.Error())
			}).
			Label("select your business case").
			ID(mainId)
		framework.RegistryNode(&node)
	}

	if err := framework.Go(mainId); err != nil {
		panic(err)
	}
}
