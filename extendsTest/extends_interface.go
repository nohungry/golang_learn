package main

import "fmt"

type Engine interface {
	Run()
	Stop()
}

type Bus struct {
	Engine
}

func (b *Bus) Run() {
	fmt.Println("Bus is running")
}

func main() {
	bus := &Bus{}
	bus.Run()
}
