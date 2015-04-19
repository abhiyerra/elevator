package main_test

import (
	. "github.com/abhiyerra/elevator"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Elevator", func() {
	var (
		elevator *Elevator
	)

	Describe("simple elevator queue", func() {
		BeforeEach(func() {
			elevator = &Elevator{
				CurrentFloor:     0,
				CurrentDirection: Up,
				NumFloors:        2,
			}

		})

		It("should move up when step is called", func() {
			elevator.AddRequestFloor(1, Down)

			// Make sure the pickup floor is enqueued.
			Expect(len(elevator.Queue)).To(Equal(1))

			Expect(elevator.Queue[0].Floor).To(Equal(1))
			Expect(elevator.Queue[0].Direction).To(Equal(Down))
			Expect(elevator.CurrentFloor).To(Equal(0))

			elevator.Step()

			// The first step should not change the
			// floor. It just changed the direction.
			Expect(len(elevator.Queue)).To(Equal(1))
			Expect(elevator.CurrentFloor).To(Equal(0))
			Expect(elevator.CurrentDirection).To(Equal(Down))

			elevator.Step()

			// This should move the elevator up to floor 1. Since
			// that is where we are picking up it should remove
			// the floor from the queue.
			Expect(elevator.CurrentFloor).To(Equal(1))
			Expect(elevator.CurrentDirection).To(Equal(Down))
			Expect(elevator.CurrentGoalFloor).To(Equal(1))
			Expect(len(elevator.Queue)).To(Equal(0))
		})
	})

	Describe("pickup for both up and down on one floor", func() {
		BeforeEach(func() {
			elevator = &Elevator{
				CurrentFloor:     0,
				CurrentDirection: Up,
				NumFloors:        3,
			}
		})

		It("goes to the correct floors", func() {
			elevator.AddRequestFloor(1, Down)
			elevator.AddRequestFloor(1, Up)

			// Make sure the pickup floors is enqueued.
			Expect(len(elevator.Queue)).To(Equal(2))
			Expect(elevator.Queue[0].Floor).To(Equal(1))
			Expect(elevator.Queue[0].Direction).To(Equal(Down))
			Expect(elevator.Queue[1].Floor).To(Equal(1))
			Expect(elevator.Queue[1].Direction).To(Equal(Up))

			elevator.AddDestinationFloor(1, 2)
			elevator.AddDestinationFloor(1, 0)

			Expect(len(elevator.DestinationQueue)).To(Equal(2))
			Expect(elevator.DestinationQueue[0].Floor).To(Equal(1))
			Expect(elevator.DestinationQueue[0].GoalFloor).To(Equal(2))
			Expect(elevator.DestinationQueue[0].Direction).To(Equal(Up))
			Expect(elevator.DestinationQueue[1].Floor).To(Equal(1))
			Expect(elevator.DestinationQueue[1].GoalFloor).To(Equal(0))
			Expect(elevator.DestinationQueue[1].Direction).To(Equal(Down))

			elevator.Step()

			// This should move the elevator up to floor 1. Since
			// that is where we are picking up it should remove
			// the floor from the queue. But it should continue
			// moving up to the destination floor of 2.
			Expect(len(elevator.Queue)).To(Equal(1))
			Expect(len(elevator.DestinationQueue)).To(Equal(2))
			Expect(elevator.CurrentFloor).To(Equal(1))
			Expect(elevator.Queue[0].Floor).To(Equal(1))
			Expect(elevator.Queue[0].Direction).To(Equal(Down))
			Expect(elevator.CurrentDirection).To(Equal(Up))

			elevator.Step()

			// We will be looking at the destination list now and
			// moving toward the destination floor of 2 and change
			// direction.
			// Expect(len(elevator.Queue)).To(Equal(1))
			// Expect(len(elevator.DestinationQueue)).To(Equal(1))
			// Expect(elevator.CurrentFloor).To(Equal(2))
			// Expect(elevator.Queue[0].Floor).To(Equal(1))
			// Expect(elevator.Queue[0].Direction).To(Equal(Down))
			// Expect(elevator.CurrentDirection).To(Equal(Down))

			// elevator.Step()

			// Go back down to floor 1. Do a pickup
			// Expect(len(elevator.Queue)).To(Equal(0))
			// Expect(len(elevator.DestinationQueue)).To(Equal(1))
			// Expect(elevator.CurrentFloor).To(Equal(1))
			// Expect(elevator.CurrentDirection).To(Equal(Down))

			// elevator.Step()

			// // Go to floor 0.
			// Expect(len(elevator.Queue)).To(Equal(0))
			// Expect(len(elevator.DestinationQueue)).To(Equal(0))
			// Expect(elevator.CurrentFloor).To(Equal(0))
			// Expect(elevator.CurrentDirection).To(Equal(Down))

			// elevator.Step()

			// // Result in a no-op
			// Expect(len(elevator.Queue)).To(Equal(0))
			// Expect(len(elevator.DestinationQueue)).To(Equal(0))
			// Expect(elevator.CurrentFloor).To(Equal(0))
			// Expect(elevator.CurrentDirection).To(Equal(Down))
		})
	})
})
