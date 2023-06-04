package main

type Multiplier interface {
	Multiply(num int) int
}

type MultiplyByX struct {
	x int
}

func (multiplier MultiplyByX) Multiply(num int) int {
	return num * multiplier.x
}