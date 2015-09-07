package elevator
import (
	"testing"
)


func TestNewGoalQueue(t *testing.T) {
	pq := NewPriorityQueue()

	if pq.Len() != 0 {
		t.Error("After initialization queue length should be 0, but was:", pq.Len())
	}
}

func TestGoalQueueAddOneItem(t *testing.T) {

	pq := NewPriorityQueue()

	pq.PushGoal(NewGoal(1,5))

	if pq.Len() == 0 {
		t.Error("After 1 x Push() queue length should be 1, but was:", pq.Len())
	}
}

func TestGoalQueuePeek(t *testing.T) {

	pq := NewPriorityQueue()

	pq.PushGoal(NewGoal(1,5))
	item := pq.Peek()

	if item == nil {
		t.Errorf("After 1 x Peek(), item should not be nil, but was")
	}

	if pq.Len() != 1 {
		t.Error("After 1 x Peek() queue length should be 1, but was:", pq.Len())
	}
}

func TestGoalQueueOrdering(t *testing.T) {


	pq := NewPriorityQueue()

	pq.PushGoal(NewGoal(1, 15))
	pq.PushGoal(NewGoal(1,5))
	pq.PushGoal(NewGoal(1,8))


	goal := pq.Pop().(*Goal)

	if goal.priority != 15 {
		t.Error("n-1 th element should have priority 15, but was:", goal.priority)
	}

	if pq.Len() != 2{
		t.Error("After 1 x Pop() queue length should be 2, but was:", pq.Len())
	}

	goal = pq.Pop().(*Goal)

	if goal.priority != 8 {
		t.Error("n-1 th element should have priority 8, but was:", goal.priority)
	}

	goal = pq.Pop().(*Goal)

	if goal.priority != 5 {
		t.Error("n-1 th element should have priority 5, but was:", goal.priority)
	}

	if pq.Len() != 0{
		t.Error("After 3 x Pop() queue length should be 0, but was:", pq.Len())
	}
}

func TestFindValueInQueuePresent(t *testing.T) {

	pq := NewPriorityQueue()
	goalToFind := NewGoal(9, 1.7)

	pq.PushGoal(NewGoal(5, 2.5))
	pq.PushGoal(NewGoal(9, 1.7))
	pq.PushGoal(NewGoal(11, 0.5))

	actualGoal := pq.Find(goalToFind.floor)

	if actualGoal == nil {
		t.Error("expectedGoal must not be nil. but was")
	}

	if actualGoal.floor != goalToFind.floor  {
		t.Errorf(" not be nil. but was: %v", actualGoal.priority)
	}

	if actualGoal.priority != 1.7 {
		t.Errorf("expectedGoal must not be nil. but was: %v", actualGoal.priority)
	}

}

func TestFindValueInQueueAbsent(t *testing.T) {

	pq := NewPriorityQueue()
	goalToFind := NewGoal(9, 1.7)

	pq.PushGoal(NewGoal(5, 2.5))
	pq.PushGoal(NewGoal(10, 1.7))
	pq.PushGoal(NewGoal(11, 0.5))

	actualGoal := pq.Find(goalToFind.floor)

	if actualGoal != nil {
		t.Errorf("expectedGoal should be nil, but was: %v", actualGoal)
	}
}

func TestFindValueInEmptyQueue(t *testing.T) {

	pq := NewPriorityQueue()
	goalToFind := NewGoal(9, 1.7)


	actualGoal := pq.Find(goalToFind.floor)

	if actualGoal != nil {
		t.Errorf("expectedGoal should be nil, but was: %v", actualGoal)
	}
}