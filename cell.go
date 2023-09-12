package main

type Cell struct {
	status byte
}

func (c Cell) IsAlive() bool {
	return c.status&1 == 1
}

func (c *Cell) Alive() {
	c.status |= 0b10
}

func (c *Cell) Kill() {
	c.status &= 0b01
}

func (c *Cell) Step() {
	c.status >>= 1
}
