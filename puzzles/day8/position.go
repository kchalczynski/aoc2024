package day8

type Position struct {
	row, col              int
	antennaType           string
	isAntenna, isAntinode bool
}

type SimplePosition struct {
	row, col int
}
