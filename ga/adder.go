package main

type Adder interface {
	Add(num int) int
}

type AddX struct {
	x int
}

func (adder AddX) Add(num int) int {
	return num + adder.x
}

type CustomAdder struct {
	x int
}