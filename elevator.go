package main

import (
	"sort"
)

type Direction int

const (
	Up Direction = iota
	Down
	None
)

type RequestType int

const (
	PickupRequest RequestType = iota
	DestinationRequest
)

type QueuedFloor struct {
	Type      RequestType
	Floor     int
	Direction Direction
}

type QueuedFloors []QueuedFloor

// For use with sort
func (qf QueuedFloors) Len() int {
	return len(qf)
}

// For use with sort
func (s QueuedFloors) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// For use with sort
func (s QueuedFloors) Less(i, j int) bool {
	return s[i].Floor < s[j].Floor
}

type Elevator struct {
	CurrentFloor     int
	CurrentDirection Direction
	Queue            QueuedFloors
}

func (e *Elevator) AddFloor(floor int, direction Direction, queueType RequestType) {
	e.Queue = append(e.Queue, QueuedFloor{
		Type:      queueType,
		Floor:     floor,
		Direction: direction,
	})
	sort.Sort(e.Queue)
}

func (e *Elevator) Step() {

	// Find all the floors greater than the current floor.
	// Go toward the nearest one going up. Optimize for going up.
	// After we've readed the top then go down. Don't change direction till we have gone to the bottom.
	// When pickedup
	// Sort them
	// Go to the first floor.

}

type ElevatorControlSystem struct {
	Elevators    []Elevator
	NumFloors    int
	NextElevator int
}

type ElevatorStatus struct {
	ElevatorId int
	Floor      int
	GoalFloor  int
}

// func (ecs *ElevatorControlSystem) Status() []ElevatorStatus {

// }

// func (ecs *ElevatorControlSystem) Update(elevatorId, floor, goalFloor int) {

// }

func (ecs *ElevatorControlSystem) Pickup(floor int, direction Direction) {
	if floor > ecs.NumFloors {
		return // That floor doesn't exist.
	}

	ecs.Elevators[ecs.NextElevator].AddFloor(floor, direction, PickupRequest)

	// A round robin way of queuing elevators. Would make it
	// distribute evently across all the elevators. Also protects
	// against bottlenecks where if we pick the nearest elevator
	// then we might be overloading that elevator with a bunch of
	// other requests.
	ecs.NextElevator++
	ecs.NextElevator = ecs.NextElevator % len(ecs.Elevators)
}

func (ecs *ElevatorControlSystem) Step() {
	for i := range ecs.Elevators {
		ecs.Elevators[i].Step()
	}
}

func NewElevatorControlSystem(numElevators, numFloors int) (ecs *ElevatorControlSystem) {
	ecs = &ElevatorControlSystem{}

	// Initialize all the elevators
	for i := 0; i < numElevators; i++ {
		ecs.Elevators = append(ecs.Elevators, Elevator{CurrentFloor: 0})
	}
	ecs.NumFloors = numFloors

	return
}

// func main() {
// 	simpleElevator := NewElevatorControl(1)

// 	simpleElevator.Pickup()
// }
