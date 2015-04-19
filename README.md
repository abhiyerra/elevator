# Elevator Control System

The elevator control system gives the entire control of pickup and
destinations to the Elevator struct. The Elevator Control System
basically uses a simple round robin algorithm to queue pickup requests
to the next increment of the elevator that is available.

## Change to the API

  - Status(): []ElevatorStatus
    - Status returns a list of ElevatorStatus structs with keys. Just
      returning a list of integers would not make a good api contract
      since it does not really specify what we are returning which is
      prone to error from the users of the API. I used a struct here
      but a hash of values would have also worked.
 - Pickup(int, Direction): int
    - Pickup returns the elevatorId since we will need to use it for
      the update. It doesn't seem to make much sense to just add a
      pickup but not knowing which elevator got the pickup
      request.
 - Update(Int, Int, Int):
    - This just queues the requests that are received to a
      queue. These are handled when the elevator does the pickup. I
      suppose this doesn't handle the case where elevators have the
      destination floor outside since it assumes you go to a pickup
      floor to choose the destination.
 - Step():
    - No change.

# Additional Changes I would make

 - We are not deduplicating requests here.
 - Right now we are just round robining all requests. Say there is a
   pickup on floor 3 for down twice. Then two elevators would both get
   those requests. Then again this is what happens in my building so
   maybe it's not a bug but a way to request an additional elevator if
   say the one that was requested gets bottlenecked.
 - Another nice scheduling algorithm would be to check the elevators
   for the length of their queues and pick the elevator with the least
   entries in its queue.
 - There would be another limitation of building where the elevators
   may not go to all the floors. The code would have to change to
   handle that.

## Building & Running

All the simulation files are in main.go. This being a go project
please follow the instructions to setup the go environment. The code
has been tested under go1.4.2. Instructions on setting up a go
environment can be gleaned here:
https://golang.org/doc/code.html#GOPATH

A simulation is provided in main.go

```
cd $GOPATH
go get -u github.com/abhiyerra/elevator
go test -cover #  To run the tests. Cover shows the test coverage.
go build .
./elevator # Should run a provided simulation.
```
