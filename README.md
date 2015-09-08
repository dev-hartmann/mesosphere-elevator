# mesosphere-elevator

For the mesos take home challenge I implemented an elevatorontrolsystem in Go

## Build Instructions

Just run go run main.go ;)

## Requirements :

* Nicer API that should Support: Pickup, Step for simulation and Status
* Better scheduling than First Come First Serve
* Elevator count up to 16 

## Implementation

General design is the following:

Two main components: Elevator and Control System

Elevators have a priority queue, with their goals(implementation of priority queue is a modified version of this golang example [Priority Queue Example](http://golang.org/pkg/container/heap/))
In each step of the simulation the elevators move towards their goals, if there are any, or otherwise are idle and waiting for new goals

The control system scores incoming pickup requests and sends them off to the most suitable elevator.

My first thought was to schedule requests based on a score that was calculated by assigning weights to the distance between elevator position and pickup floor.
After some research I found a Nearest Car/Elevator Algorithm as described in these two papers([Paper 1](http://www.columbia.edu/~cs2035/courses/ieor4405.S13/p14.pdf), [Paper 2](http://www.diva-portal.org/smash/get/diva2:811554/FULLTEXT01.pdf)).


```go
func calculateSuitabilityScore(numberOfFloors int, elevator *Elevator, req *Request) float64 {

	//if requested floor is already inside goal queue, return its priority
	if goal := elevator.Goals.Find(req.Floor); goal != nil {
		return goal.priority
	}

	direction := elevator.Direction
	distance := elevator.Distance(req)

	//if elevator is moving towards goal 
	if direction * distance >= 0 {
		//in same direction
		if direction == req.Direction {
			return float64(numberOfFloors+2) - math.Abs(float64(distance))
		}
		//different direction or is idle
		return float64(numberOfFloors+1) - math.Abs(float64(distance))
	}
	//elevator is away from goal
	return 1.0
}
´´´

Due to time limitations comments and discussion about scheduling algorithm are a bit shorter than it should be.