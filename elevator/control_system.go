package elevator

import (
	"math"
)

type ControlSystem struct {
	Elevators         []*Elevator
	NumberOfElevators int
	NumberOfFloors    int
}

type ControlSystemStatus []*Elevator

func(css *ControlSystemStatus)Len() int {
	return len(*css)
}

func(css *ControlSystemStatus)GetStatusAtIndex(index int) *Elevator {
	cpy := *css
	return cpy[index]
}

func NewConstrolSystem(numOfElevators, maxFloors int) *ControlSystem {

	cs := &ControlSystem{
		NumberOfElevators: numOfElevators,
		NumberOfFloors:    maxFloors,
	}

	for i := 0; i < numOfElevators; i++ {
		elv := NewElevator(i, 0)
		cs.Elevators = append(cs.Elevators, elv)
	}

	return cs
}

func calculateSuitabilityScore(numberOfFloors int, elevator *Elevator, req *Request) float64 {

	if goal := elevator.Goals.Find(req.Floor); goal != nil {
		return goal.priority
	}

	direction := elevator.Direction
	distance := elevator.Distance(req)

	if direction * distance >= 0 {
		if direction == req.Direction {
			return float64(numberOfFloors+2) - math.Abs(float64(distance))
		}
		return float64(numberOfFloors+1) - math.Abs(float64(distance))
	}
	return 1.0
}

func (cs *ControlSystem) PickUp(floor, direction int) {
	req := NewRequest(floor, direction)

	highestScore := 0.0
	var chosenElevator *Elevator

	for _, elevator := range cs.Elevators {
		score := calculateSuitabilityScore(cs.NumberOfFloors, elevator, req)
		if score > highestScore {
			highestScore = score
			chosenElevator = elevator
		}
	}
	chosenElevator.AddGoal(req, highestScore)
}

func (cs *ControlSystem) Status() *ControlSystemStatus{
	css := ControlSystemStatus(cs.Elevators)
	return &css
}

func (cs *ControlSystem) Step(steps int) {
	for i := 0; i < steps; i++ {
		for _, elevator := range cs.Elevators {
			elevator.Step()
		}
	}
}
