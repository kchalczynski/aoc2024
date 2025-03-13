package day10

type Position struct {
	// row and col
	y, x int
	// numerical value
	height int
	// how many positions can be visited from that one
	score int
}
