package main

import (
	elv "github.com/Elevator-System/elevator"
)

func main() {

	c := elv.NewConstrolSystem(2, 10)

	c.PickUp(3, elv.UP)
	c.PickUp(2, elv.UP)
	c.PickUp(5, elv.UP)
	c.PickUp(0, elv.UP)
	c.PickUp(2, elv.DOWN)
	c.PickUp(6, elv.UP)

	for i := 0; i < 7; i++ {
		c.Step(1)
	}

	c.PickUp(1, elv.DOWN)
	c.PickUp(5, elv.DOWN)
	c.PickUp(2, elv.DOWN)
	c.PickUp(6, elv.UP)
	c.PickUp(9, elv.DOWN)
	c.PickUp(7, elv.UP)
	c.PickUp(1, elv.DOWN)
	c.PickUp(0, elv.DOWN)


	for i := 0; i < 9; i++ {
		c.Step(1)
	}
}
