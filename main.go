package main

import (
	"fmt"
)

func Simulation1() {
	ecs := NewControlSystem(3, 10)

	ecs.Pickup(5, Up)
	elevatorId := ecs.Pickup(2, Up)
	ecs.Update(elevatorId, 2, 6)
	ecs.Pickup(3, Up)
	ecs.Pickup(3, Down)

	for i := 0; i < 10; i++ {
		ecs.Step()
		fmt.Printf("Step %d:\n%v\n\n", i, ecs.Status())
	}
}

func main() {
	Simulation1()
}
