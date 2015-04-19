package main_test

import (
	. "github.com/abhiyerra/elevator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ControlSystem", func() {
	PContext("with 2 floors", func() {
		var (
			ecs *ControlSystem
		)

		BeforeEach(func() {
			ecs = NewControlSystem(1, 2)
		})

		It("should have the correct initial fields", func() {
			Expect(len(ecs.Elevators)).To(Equal(1))
			Expect(ecs.NumFloors).To(Equal(2))

			for _, i := range ecs.Elevators {
				Expect(i.CurrentFloor).To(Equal(0))
			}
		})

		It("should not queue invalid requests", func() {
			ecs.Pickup(3, Down)
			Expect(len(ecs.Elevators[0].Queue)).To(Equal(0))
		})

		It("should queue requests ito the elevator's queue", func() {
			ecs.Pickup(2, Down)
			Expect(len(ecs.Elevators[0].Queue)).To(Equal(1))
			ecs.Step()
			//			Expect(ecs.Elevators[0].CurrentFloor).To(Equal(2))
		})

		// Request Pickup on Floor2 to go down
		// Step -> Move elevator up one floor.
		// Step -> Pick up person
		// Step -> Move elevator down one floor.
	})

	PContext("With 3 floors", func() {
		// Floor 1
		// Request Pickup on Floor2 to go down to Floor1
		// Request Pickup on Floor5 to go up to Floor 9
		// Request Pickup on Floor3 to go down to Floor 2
		// Request Pickup on Floor4 to go up to Floor 6
		// Step -> Move elevator up one floor.
		// Step -> Pick up person
		// Step -> Move elevator down one floor.
	})

	Describe("Pickup", func() {
		var (
			ecs *ControlSystem
		)

		BeforeEach(func() {
			ecs = NewControlSystem(3, 3)
		})

		It("should add floors in a round robin", func() {
			Expect(len(ecs.Elevators)).To(Equal(3))
			Expect(ecs.NumFloors).To(Equal(3))

			ecs.Pickup(1, Up)
			Expect(len(ecs.Elevators[0].Queue)).To(Equal(1))
			Expect(len(ecs.Elevators[1].Queue)).To(Equal(0))
			Expect(len(ecs.Elevators[2].Queue)).To(Equal(0))
			Expect(ecs.CurrentElevator).To(Equal(1))

			ecs.Pickup(2, Down)
			Expect(len(ecs.Elevators[0].Queue)).To(Equal(1))
			Expect(len(ecs.Elevators[1].Queue)).To(Equal(1))
			Expect(len(ecs.Elevators[2].Queue)).To(Equal(0))
			Expect(ecs.CurrentElevator).To(Equal(2))

			ecs.Pickup(3, Up)
			Expect(len(ecs.Elevators[0].Queue)).To(Equal(1))
			Expect(len(ecs.Elevators[1].Queue)).To(Equal(1))
			Expect(len(ecs.Elevators[2].Queue)).To(Equal(1))
			Expect(ecs.CurrentElevator).To(Equal(0))

			ecs.Pickup(3, Down)
			Expect(len(ecs.Elevators[0].Queue)).To(Equal(2))
			Expect(len(ecs.Elevators[1].Queue)).To(Equal(1))
			Expect(len(ecs.Elevators[2].Queue)).To(Equal(1))
			Expect(ecs.CurrentElevator).To(Equal(1))

			ecs.Pickup(200, Up) // Doesn't exist.
			Expect(len(ecs.Elevators[0].Queue)).To(Equal(2))
			Expect(len(ecs.Elevators[1].Queue)).To(Equal(1))
			Expect(len(ecs.Elevators[2].Queue)).To(Equal(1))
			Expect(ecs.CurrentElevator).To(Equal(1))
		})

	})
})
