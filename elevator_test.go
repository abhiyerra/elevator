package main_test

import (
	"fmt"
	. "github.com/abhiyerra/elevator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Elevator", func() {
	var (
		elevator *Elevator
	)

	BeforeEach(func() {
		elevator = &Elevator{
			CurrentFloor:     0,
			CurrentDirection: Up,
		}
	})

	It("should move up when step is called", func() {
		elevator.AddRequestFloor(1, Down)

		Expect(len(elevator.Queue)).To(Equal(1))
		fmt.Println(elevator)
		Expect(elevator.Queue[0].Floor).To(Equal(1))
		Expect(elevator.Queue[0].Direction).To(Equal(Down))

		elevator.Step()

		fmt.Println(elevator)

		Expect(len(elevator.Queue)).To(Equal(1))
		Expect(elevator.CurrentFloor).To(Equal(1))

		// elevator.Step()

		// Expect(len(elevator.Queue)).To(Equal(0))
		// Expect(elevator.CurrentFloor).To(Equal(1))
	})

})
