package main_test

import (
	. "github.com/abhiyerra/elevator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ControlSystem", func() {
	var (
		ecs *ControlSystem
	)

	Describe("NewControlSystem", func() {
		BeforeEach(func() {
			ecs = NewControlSystem(1, 2)
		})

		It("should have the correct initial fields", func() {
			Expect(len(ecs.Elevators)).To(Equal(1))
			Expect(ecs.NumFloors).To(Equal(2))
			Expect(ecs.CurrentElevator).To(Equal(0))

			for _, i := range ecs.Elevators {
				Expect(i.CurrentFloor).To(Equal(0))
			}
		})

	})

	Describe("Pickup", func() {
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

	Describe("Update", func() {
		var (
			elevatorId int
		)

		BeforeEach(func() {
			ecs = NewControlSystem(3, 10)
			elevatorId = ecs.Pickup(4, Up)
		})

		It("should add the destination floor", func() {
			ecs.Update(elevatorId, 4, 7)

			Expect(len(ecs.Elevators[0].Queue)).To(Equal(1))
			Expect(len(ecs.Elevators[0].DestinationQueue)).To(Equal(1))

			Expect(len(ecs.Elevators[1].Queue)).To(Equal(0))
			Expect(len(ecs.Elevators[2].Queue)).To(Equal(0))
			Expect(ecs.CurrentElevator).To(Equal(1))

			ecsStatus := ecs.Status()
			Expect(len(ecsStatus)).To(Equal(3))
			Expect(ecsStatus[0].ElevatorId).To(Equal(0))
			Expect(ecsStatus[0].Floor).To(Equal(1))
			Expect(ecsStatus[0].GoalFloor).To(Equal(1))

			ecs.Step()

			ecsStatus = ecs.Status()
			Expect(ecsStatus[0].ElevatorId).To(Equal(0))
			Expect(ecsStatus[0].Floor).To(Equal(2))
			Expect(ecsStatus[0].GoalFloor).To(Equal(4))

			ecs.Step()
			ecsStatus = ecs.Status()
			Expect(ecsStatus[0].ElevatorId).To(Equal(0))
			Expect(ecsStatus[0].Floor).To(Equal(3))
			Expect(ecsStatus[0].GoalFloor).To(Equal(4))
		})
	})

})
