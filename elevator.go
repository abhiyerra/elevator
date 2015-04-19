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
	log.Println("queuing floor", floor)
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

	e.Queue = append(e.DestinationQueue, QueuedFloor{
		Floor:     floor,
		Direction: direction,
		GoalFloor: goalFloor,
	})

	sort.Sort(e.DestinationQueue)
}

func (e *Elevator) goTowards(queuedFloor *QueuedFloor) {
	if queuedFloor.Floor > e.CurrentFloor {
		e.CurrentFloor++
	} else {
		e.CurrentFloor--
	}
}

func (e *Elevator) Step() {
	if e.Queue.Len() == 0 {
		log.Println("No more steps to take")
		return
	}

	var goingTowards QueuedFloor

	if e.CurrentDirection == Up {
		var upwardQueue QueuedFloors

		log.Println("Going up")

		for i := 0; i < len(e.Queue); i++ {
			log.Println("Going up, to floor", i, "current floor", e.CurrentFloor)

			if e.Queue[i].Direction == Up && e.Queue[i].Floor >= e.CurrentFloor {
				log.Println("Going up, to floor", e.Queue[i].Floor)

				upwardQueue = append(upwardQueue, e.Queue[i])
			}
		}

		if upwardQueue.Len() == 0 && e.Queue.Len() > 0 {
			log.Println("No where to go up, going down")
			e.CurrentDirection = Down
		} else {
			sort.Sort(upwardQueue)

			upwardToFloor := upwardQueue[0]
			log.Println("Going up to floor", upwardToFloor)
			e.goTowards(&upwardToFloor)
		}
	}

	if e.CurrentDirection == Down {
		var downwardQueue QueuedFloors

		for i := 0; i < len(e.Queue); i++ {
			if e.Queue[i].Direction == Up && e.Queue[i].Floor < e.CurrentFloor {
				downwardQueue = append(downwardQueue, e.Queue[i])
			}
		}

		if downwardQueue.Len() == 0 && e.Queue.Len() > 0 {
			e.CurrentDirection = Up
		} else {
			sort.Sort(downwardQueue)

			e.goTowards(&downwardQueue[0])
		}
	}

	e.CurrentGoalFloor = goingTowards.Floor

	// if we are at that floor. Remove the item from the floorqueue.
}
