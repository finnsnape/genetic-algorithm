package main

import (
	"github.com/traefik/yaegi/interp"
)

func runAnswer(newAnswer answer) bool {
	i := interp.New(interp.Options{})

	_, err := i.Eval(newAnswer.Code)
	if err != nil {
		return false
	}

	v, err := i.Eval("Bar")
	if err != nil {
		return false
	}

	bar := v.Interface().(func(int, int) bool)

	r := bar(2, 4)
	println(r)

	return true
}
