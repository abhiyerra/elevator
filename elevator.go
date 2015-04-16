package main

type Direction int

const (
	Up Direction = iota
	Down
)

type Elevator struct {
	CurrentFloor int
	Queue        []int
}

type ElevatorControlSystem struct {
	Elevators []Elevator
	NumFloors int
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

// func (ecs *ElevatorControlSystem) Pickup(floor int, direction Direction) {

// }

// func (ecs *ElevatorControlSystem) Step() {
// }

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
