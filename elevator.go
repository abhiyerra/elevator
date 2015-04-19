package main

import (
	"sort"
)

type Direction int

const (
	Up Direction = iota
	Down
)

type Elevator struct {
	CurrentFloor     int // Starts at 0
	CurrentDirection Direction
	Queue            QueuedFloors
	DestinationQueue QueuedFloors
	CurrentGoalFloor int
	NumFloors        int
}

func (e *Elevator) AddRequestFloor(floor int, direction Direction) {
	e.Queue = append(e.Queue, QueuedFloor{
		Floor:     floor,
		Direction: direction,
	})

	sort.Sort(e.Queue)
}

func (e *Elevator) AddDestinationFloor(floor int, goalFloor int) {
	var direction Direction
	if goalFloor > floor {
		direction = Up
	} else {
		direction = Down
	}

	e.DestinationQueue = append(e.DestinationQueue, QueuedFloor{
		Floor:     floor,
		Direction: direction,
		GoalFloor: goalFloor,
	})

	sort.Sort(e.DestinationQueue)
}

func (e *Elevator) upwardAndDownwardQueue() (upwardQueue QueuedFloors, downwardQueue QueuedFloors) {
	for i := 0; i < len(e.Queue); i++ {
		switch {
		case e.Queue[i].Direction == Up:
			upwardQueue = append(upwardQueue, e.Queue[i])
		case e.Queue[i].Direction == Down:
			downwardQueue = append(downwardQueue, e.Queue[i])
		}
	}

	sort.Sort(upwardQueue)
	sort.Sort(downwardQueue)

	return
}

func (e *Elevator) appendDestinationFloors() {
	destFloors := e.DestinationQueue.DestsWithFloorAndDirection(e.CurrentFloor, e.CurrentDirection)

	for i := 0; i < len(destFloors); i++ {
		e.Queue = append(e.Queue, QueuedFloor{
			Floor:     destFloors[i].GoalFloor,
			Direction: destFloors[i].Direction,
		})
	}

	e.DestinationQueue = e.DestinationQueue.RemoveFloorAndDirection(e.CurrentFloor, e.CurrentDirection)
}

func (e *Elevator) Step() {
	var (
		goingTowards               *QueuedFloor
		upwardQueue, downwardQueue = e.upwardAndDownwardQueue()
	)

	switch {
	case e.CurrentDirection == Up && upwardQueue.Len() == 0 && e.Queue.Len() > 0:
		e.CurrentDirection = Down
	case e.CurrentDirection == Up && upwardQueue.Len() > 0:
		goingTowards = upwardQueue.NextUpward(e.CurrentFloor)
	case e.CurrentDirection == Down && downwardQueue.Len() == 0 && e.Queue.Len() > 0:
		e.CurrentDirection = Up
	case e.CurrentDirection == Down && downwardQueue.Len() > 0:
		goingTowards = downwardQueue.NextDownward(e.CurrentFloor)
	case e.Queue.Len() == 0:
		return
	}

	if goingTowards != nil {
		e.CurrentGoalFloor = goingTowards.Floor

		if goingTowards.Floor > e.CurrentFloor {
			e.CurrentFloor++
		} else {
			e.CurrentFloor--
		}
	}

	e.Queue = e.Queue.RemoveFloorAndDirection(e.CurrentFloor, e.CurrentDirection)
	e.appendDestinationFloors()
}
