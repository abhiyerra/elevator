package main

type QueuedFloor struct {
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
