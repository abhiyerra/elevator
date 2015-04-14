package main

type Direction int

const(
	Up Direction = iota
	Down 
)

type ElevatorControlSystem interface {
	//   Status(): Seq[(Int, Int, Int)]
	//   Update(Int, Int, Int)
	Pickup(floor int, direction Direction)
	//   Step()
}

type SimpleElevator struct {
	Elevators []int
}

func NewSimpleElevator(numElevators int) (se SimpleElevator) {
	// Initialize all the elevators to
	for i := 0; i < numElevators; i++ {
		se.Elevators = append(se.Elevators, 0)
	}

	return
}

func main() {
	simpleElevator := NewSimpleElevator(1)

	simpleElevator.Pickup(

}
