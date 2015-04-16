package main_test

import (
	. "github.com/abhiyerra/elevator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Elevator", func() {
	Describe("Single Elevator", func() {
		Context("With 2 floors", func() {
			var (
				ecs *ElevatorControlSystem
			)

			BeforeEach(func() {
				ecs = NewElevatorControlSystem(1, 2)
			})

			It("should have the correct initial fields", func() {
				Expect(len(ecs.Elevators)).To(Equal(1))
				Expect(ecs.NumFloors).To(Equal(2))

				for _, i := range ecs.Elevators {
					Expect(i.CurrentFloor).To(Equal(0))
				}
			})

			//			It("should queue a
			// Floor 1
			// Request Pickup on Floor2 to go down
			// Step -> Move elevator up one floor.
			// Step -> Pick up person
			// Step -> Move elevator down one floor.
		})

		Context("With 10 floors", func() {
			// Floor 1
			// Request Pickup on Floor2 to go down to Floor1
			// Request Pickup on Floor5 to go up to Floor 9
			// Request Pickup on Floor3 to go down to Floor 2
			// Request Pickup on Floor4 to go up to Floor 6
			// Step -> Move elevator up one floor.
			// Step -> Pick up person
			// Step -> Move elevator down one floor.
		})

	})

})
