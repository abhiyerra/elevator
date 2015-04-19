package main

type QueuedFloor struct {
	Floor     int
	Direction Direction
	// For use with destination after pickup
	GoalFloor int
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

func (s QueuedFloors) RemoveFloorAndDirection(floor int, direction Direction) (newQueue QueuedFloors) {
	for i := 0; i < s.Len(); i++ {
		if s[i].Floor == floor && s[i].Direction == direction {
			// No-op. Maybe we want to track the floors we
			// were on for later optimizations?
		} else {
			newQueue = append(newQueue, s[i])
		}
	}

	return newQueue
}

func (s QueuedFloors) NextUpward(floor int) *QueuedFloor {
	return &s[0]
}

func (s QueuedFloors) NextDownward(floor int) *QueuedFloor {
	return &s[s.Len()-1]
}

func (s QueuedFloors) DestsWithFloorAndDirection(floor int, direction Direction) (q QueuedFloors) {
	for i := 0; i < s.Len(); i++ {
		if s[i].Floor == floor && s[i].Direction == direction {
			q = append(q, s[i])
		}
	}

	return q
}
