package elevator

import (
	"container/heap"
	"sort"
	"fmt"
	"strings"
)

type Goal struct {
	floor    int
	priority float64
	index    int
}

type PriorityQueue []*Goal

func NewPriorityQueue() *PriorityQueue {
	pq := new(PriorityQueue)
	heap.Init(pq)
	return pq
}

func NewGoal(goalFloor int, prio float64) *Goal {
	return &Goal{floor: goalFloor, priority: prio}
}

func(pq *PriorityQueue) PushGoal(goal *Goal) {
	heap.Push(pq,goal)
	pq.update(goal, goal.floor, goal.priority)
	sort.Sort(pq)
}

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Goal)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) update(goal *Goal, value int, priority float64) {
	goal.floor = value
	goal.priority = priority
	heap.Fix(pq, goal.index)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq PriorityQueue) Peek() *Goal {
	n := len(pq)
	goal := pq[n - 1]
	return goal
}

func (pq PriorityQueue) Find(value int) *Goal{
	for i := 0; i < pq.Len(); i++ {
		goal := pq[i]
		if goal.floor == value {
			return goal
		}
	}
	return nil
}

func (pq PriorityQueue) PrintGoals() string {

	resultString := make([]string, pq.Len())

	for i := 0; i < pq.Len(); i++ {
		goal := pq[pq.Len() - 1 - i]
		resultString[i] = fmt.Sprintf("Goal #%v (floor: %v priority: %v)",i, goal.floor, goal.priority)
	}
	return strings.Join(resultString, " ")
}