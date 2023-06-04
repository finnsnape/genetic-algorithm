package main
//
//import (
//	"fmt"
//	"github.com/traefik/yaegi/interp"
//	"reflect"
//)
//
//type Breeder interface {
//	Breed(parent1 []byte, parent2 []byte) ([]byte, []byte)
//}
//
//type SomeBreeder struct{}
//
//func (SomeBreeder) Breed(parent1 []byte, parent2 []byte) ([]byte, []byte) {
//	v, err := i.Eval(function)
//	if err != nil {
//		panic(err)
//	}
//	a := v.Interface().(func ([]byte, []byte) ([]byte, []byte))
//	return a(parent1, parent2)
//}
//
//func createBreeder(i *interp.Interpreter, function string, functionName string) Breeder {
//	a := reflect.New()
//}
//
//func getFunction(v reflect.Value, functionName string) func(any) {
//	//interf := v.Interface()
//	switch functionName {
//	case "breeder":
//		return (func([]byte, []byte) ([]byte, []byte))
//	case "mutator":
//		return func([]byte) []byte)
//	case "selector":
//		return func([]string), string// need Individual type here
//	case "evaluator":
//		return func([]byte) float64
//	case "test":
//		return func(int) int
//	}
//}
//
//func buildFunction(function string, functionName string) {
//	i := interp.New(interp.Options{})
//	v, err := i.Eval(function)
//
//	if err != nil {
//		panic(err)
//	}
//
//
//
//	newFunction := getFunction(v, functionName)
//	r := newFunction(1)
//	fmt.Println(r)
//	// can we use our existing interfaces?
//
//}
//
//func runAnswer(newAnswer answer) bool {
//	i := interp.New(interp.Options{})
//
//	_, err := i.Eval(newAnswer.Code)
//	if err != nil {
//		return false
//	}
//
//	v, err := i.Eval("Bar")
//	if err != nil {
//		return false
//	}
//	typeA := typeMap["int"]
//	bar := v.Interface().(func(typeA, int) bool)
//
//	r := bar(2, 4)
//	println(r)
//
//	return true
//}
