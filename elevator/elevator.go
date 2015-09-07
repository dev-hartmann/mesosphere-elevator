package elevator

import (
	"fmt"
)

const (
	DOWN = -1
	IDLE = 0
	UP   = 1
)

type Elevator struct {
	Id           int
	CurrentFloor int
	Goals        *PriorityQueue
	Direction    int
}


func NewElevator(id, startingFloor int) *Elevator {
	return &Elevator{
		Id:           id,
		CurrentFloor: startingFloor,
		Goals:        NewPriorityQueue(),
		Direction:    IDLE,
	}
}

func (e *Elevator) Step() {

	if e.Goals.Len() == 0 {
		e.Direction = IDLE
		return
	}

	currentGoal := e.Goals.Peek()
	e.move(currentGoal.floor)
	fmt.Printf("Elevator with ID:%v on Floor: %v with Direction: %v and Current Goal:%v  Goals: %v\n", e.Id, e.CurrentFloor, e.Direction,currentGoal, e.Goals.PrintGoals())

}

func (e *Elevator) Distance(req *Request) int {
	return req.Floor - e.CurrentFloor
}

func (e *Elevator) AddGoal(req *Request, priority float64) {

	if req.Floor == e.CurrentFloor {
		return
	}

	goal := NewGoal(req.Floor, priority)
	e.Goals.PushGoal(goal)
}

func (e *Elevator) move(goalFloor int) {

	if goalFloor > e.CurrentFloor {
		e.Direction = UP
		e.CurrentFloor++

	} else {
		e.Direction = DOWN
		e.CurrentFloor--
	}

	if goalFloor== e.CurrentFloor {
		e.Goals.Pop()
		if e.Goals.Len() == 0 {
			e.Direction = IDLE
		}
	}
}
