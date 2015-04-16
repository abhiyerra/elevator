package main

import (
	"sort"
)

type Direction int

const (
	Up Direction = iota
	Down
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

func (e *Elevator) goTowards(queuedFloor *QueuedFloor) {

}

func (e *Elevator) Step() {
	if e.Queue.Len() == 0 {
		return
	}

	var goingTowards QueuedFloor

changedDirections:
	if e.CurrentDirection == Up {
		var upwardQueue QueuedFloors

		for i := range e.Queue {
			if e.Queue[i].Direction == Up && e.Queue[i].Floor >= e.CurrentFloor {
				upwardQueue = append(upwardQueue, e.Queue[i])
			}
		}

		if upwardQueue.Len() == 0 && e.Queue.Len() > 0 {
			e.CurrentDirection = Down

			// A goto! Well I'm going to love explain this
			// to you guys. But I feel it is appropriate
			// here since we are basically restarting this
			// flow to change directions. We could also
			// hold variables and what not to force the
			// else condition to run but this just seems
			// easier. :)
			goto changedDirections
		}

		sort.Sort(upwardQueue)

		goingTowards = upwardQueue[0]
	} else {
		var downwardQueue QueuedFloors

		for i := range e.Queue {
			if e.Queue[i].Direction == Up && e.Queue[i].Floor < e.CurrentFloor {
				downwardQueue = append(downwardQueue, e.Queue[i])
			}
		}

		if downwardQueue.Len() == 0 && e.Queue.Len() > 0 {
			e.CurrentDirection = Up

			goto changedDirections
		}

		sort.Sort(downwardQueue)

		goingTowards = downwardQueue[0]
	}

	e.goTowards(&goingTowards)
	// if we are at that floor. Remove the item from the floorqueue.
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

	// Initialize all the elevators. Just default to going up
	// since we are basically stating at the ground floor.
	for i := 0; i < numElevators; i++ {
		ecs.Elevators = append(ecs.Elevators, Elevator{
			CurrentDirection: Up,
			CurrentFloor:     0,
		})
	}
	ecs.NumFloors = numFloors

	return
}

// func main() {
// 	simpleElevator := NewElevatorControl(1)

// 	simpleElevator.Pickup()
// }
