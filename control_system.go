package main

type ControlSystem struct {
	Elevators       []Elevator
	NumFloors       int
	CurrentElevator int
}

type ElevatorStatus struct {
	ElevatorId int
	Floor      int
	GoalFloor  int
}

func (ecs *ControlSystem) Status() (es []ElevatorStatus) {
	for i, e := range ecs.Elevators {
		es = append(es, ElevatorStatus{
			ElevatorId: i,
			Floor:      e.CurrentFloor,
			GoalFloor:  e.CurrentGoalFloor,
		})
	}

	return
}

// A round robin way of queuing elevators. Would make it
// distribute evently across all the elevators. Also protects
// against bottlenecks where if we pick the nearest elevator
// then we might be overloading that elevator with a bunch of
// other requests.
func (ecs *ControlSystem) nextPickupElevator() int {
	curPickup := ecs.CurrentElevator

	ecs.CurrentElevator++
	// Restart the increment.
	if ecs.CurrentElevator == len(ecs.Elevators) {
		ecs.CurrentElevator = 0
	}

	return curPickup
}

// I'm a bit confused about this method. Is it saying that when a
// Pickup happens on a floor then we need to queue in a destination
// floor? That is what this method is assuming.
func (ecs *ControlSystem) Update(elevatorId, floor, goalFloor int) {
	ecs.Elevators[elevatorId].AddDestinationFloor(floor, goalFloor)
}

func (ecs *ControlSystem) Pickup(floor int, direction Direction) (elevatorId int) {
	if floor > ecs.NumFloors {
		return // That floor doesn't exist.
	}

	elevatorId = ecs.nextPickupElevator()
	ecs.Elevators[elevatorId].AddRequestFloor(floor-1, direction)

	return elevatorId
}

func (ecs *ControlSystem) Step() {
	for i := range ecs.Elevators {
		ecs.Elevators[i].Step()
	}
}

func NewControlSystem(numElevators, numFloors int) (ecs *ControlSystem) {
	ecs = &ControlSystem{}

	// Initialize all the elevators. Just default to going up
	// since we are basically stating at the ground floor.
	for i := 0; i < numElevators; i++ {
		ecs.Elevators = append(ecs.Elevators, Elevator{
			CurrentDirection: Up,
			CurrentFloor:     0,
			NumFloors:        numFloors,
		})
	}
	ecs.NumFloors = numFloors
	ecs.CurrentElevator = 0

	return
}

func main() {

}
