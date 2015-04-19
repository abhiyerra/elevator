package main

import (
	"log"
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

	log.Println("queuing floor", floor, goalFloor, direction == Up)

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

func (e *Elevator) Step() {
	var (
		goingTowards               *QueuedFloor
		upwardQueue, downwardQueue = e.upwardAndDownwardQueue()
	)

	log.Println("Floors above", e.CurrentFloor, len(upwardQueue), e.Queue[0].Direction == Up)
	log.Println("Floors below", e.CurrentFloor, len(downwardQueue))

	switch {
	case e.CurrentDirection == Up && upwardQueue.Len() == 0 && e.Queue.Len() > 0:
		log.Println("No where to go up, going down")
		e.CurrentDirection = Down
	case e.CurrentDirection == Up && upwardQueue.Len() > 0:
		goingTowards = upwardQueue.NextUpward(e.CurrentFloor)
		log.Println("Going up to floor", goingTowards)
	case e.CurrentDirection == Down && downwardQueue.Len() == 0 && e.Queue.Len() > 0:
		log.Println("No where to go down, going up")
		e.CurrentDirection = Up
	case e.CurrentDirection == Down && downwardQueue.Len() > 0:
		goingTowards = downwardQueue.NextDownward(e.CurrentFloor)
		log.Println("Going down to floor", goingTowards)
	case e.Queue.Len() == 0:
		log.Println("No more steps to take")
		return
	}

	if goingTowards != nil {
		e.CurrentGoalFloor = goingTowards.Floor
		log.Println("Going towards", goingTowards, "current goal floor", e.CurrentGoalFloor)

		if goingTowards.Floor > e.CurrentFloor {
			e.CurrentFloor++
		} else {
			e.CurrentFloor--
		}
	}

	e.Queue = e.Queue.RemoveFloorAndDirection(e.CurrentFloor, e.CurrentDirection)
}
