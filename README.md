# Elevator Control System

Things to be aware of:

 - [ ] How many elevators?
 - [ ] How many people are waiting for the elevator in each floor?
 - [ ] How fast does the elevator go up and down?  Is the speed constant
   per floor or does it have one of those accelerators?
 - [ ] Can we upgrade the system at a later time? Can we updating the
   scheduling system?
 - [ ] What is the rest state of the elevators? Do they all go down or
   do they spread out so there can be an optimal time to pick up an
   elevator?
 - [ ] How do we handle the numbering of the elevators?

Test cases:
 - [ ] What's the case for one elevator?
 - [ ] What's the case for 2 elevators?
 - [ ] What's the case for n elevators?
 - [ ] What happens if every button is pressed?
 - [ ] What happens if there are people on various floors?

Changes to the API:

 - Pickup(floor int, direction Direction): I will be using a Direction
   type so we have an abstraction to represent the direction.
 - Have an elevator type. Where we queue up the various floors an
   elevator should visit.


There are multiple scheduling algorithms based on different aspects of
the elevators, the building, etc.

 - The design I use is simple in that it that the elevators go in a
   single direction based on the queuing. So an elevator would have to
   know what direction it is going and queue in all the






Why do it this way? We can give the elevator a level of responsibility
and distribute the load throughout the building. So if there are a bunch of people


Other algorithms to use:

 - If say in a 10 floor building the 7th floor paid a premium to
   always receive the elevators first then we would use a priority
   queue for queuing requests.
 - If we have a

 - Say we have 3 elevators and 30 floors. The we divide the elevators
 into the floors to get 10 floors that each elevator is responsible
 for.
   - So elevator 1 is responsible for floors 1-10
   - So elevator 2 is responsible for floors 1-20
   - So elevator 3 is responsible for floors 1-30
 - They are responsible for them insofar as requests are made on those
   floors. The elevators themselves can go to whatever floors.
 - When a person requests a pickup an elevator designated for that
   floor picks them up.
    - The pick up happens in a list with the highest requested at the
      top.
    - When the pickup is done and there are no more people in the
      elevator it goes back to the first floor.
