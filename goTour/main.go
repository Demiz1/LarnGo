package main

import (
	"errors"
	"fmt"
	"time"
)

func simpleAdd(x int, y int) int {
	return x + y
}
func twoReturn(x string) (string, string) {
	return x, x
}

func variableTests() {
	fmt.Println("Hello world!", time.Now())
	a := 1
	b := 1336
	fmt.Println("now i add two values", simpleAdd(a, b))
	s := "hello"
	var ost string
	_, ost = twoReturn(s)
	fmt.Println("here i use a function with two returns", ost)

	v := 42 // change me!
	fmt.Printf("v is of type %T\n", uint16(v))
	v_2 := float64(v)
	fmt.Printf("%T\n", v_2)

	var g, j, c float64 = 1, 24, 5

	fmt.Printf("%T,%T,%T\n", g, j, c)
}

func forLoopTests() {
	fmt.Println("==========LOOP LEARNING=============================")
	// sum 10 numbers
	x := 0
	for i := 0; i < 10; i++ {
		x += 1
	}
	fmt.Printf("after loop, x is %d\n", x)

	// while loop is defined with the first and last "for"-arguments ommited
	var gurk, ost int = 6, 0
	for gurk < 278 {
		gurk *= 2
		ost += 1
	}
	fmt.Printf("we reached number %d, we needed  %d loops \n", gurk, ost)

	// a better implementation perhaps idk
	gurk = 6
	for ost = 0; gurk < 278; ost++ {
		gurk *= 2
	}
	fmt.Printf("we reached number %d, we needed  %d loops \n", gurk, ost)

}

func ifTests() error {
	a, b := 5, 20

	if a < b {
		fmt.Printf("Yes, a is less than b, %d,%d\n", a, b)
	} else {
		return errors.New("this is very odd")
	}

	if g := 5; g < 20 {
		fmt.Printf("you can apparently create variables in the if-statement itself\n")
	}

	//just a quick test
	h := "a"
	fmt.Printf("the numerical value of \"a\" is %d\n", int(h[0]))
	return nil
}

func switchTest() {
	fmt.Printf("==============Switch======================\n")
	gaga := "hello"
	switch gaga {
	case "hi":
		fmt.Printf("The text is gaga\n")
	case "hello":
		fmt.Printf("The value is \"hello\"\n")
	default:
		fmt.Printf("there was no matching case...\n")
	}
}
func main() {
	variableTests()
	forLoopTests()
	ifTests()
	switchTest()
}
