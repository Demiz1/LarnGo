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

func pointerTest() {
	fmt.Printf("==============Pointers======================\n")
	g := 1337.5
	p := &g
	// more explicit pointer extraction
	var p_1 *float64 = &g

	fmt.Printf("The value g:%f have the pointer %d\n", g, p)
	if p == p_1 {
		fmt.Printf("Yes, both methods gave the same address\n")
	}

	// we can also ofc change values through pointer deferal
	*p = 55
	fmt.Printf("now the variable g have the value %d\n", uint64(*p))
}

func structTest() {
	fmt.Printf("==============STRUCT======================\n")
	// a struct is created like this
	type myStruct struct {
		a int
		b float64
		c string
	}
	// instance
	a := myStruct{1, 15.8982384, "hello"}
	b := myStruct{}

	// neat we can print a whole struct with the print function
	fmt.Println(a, b)

	b.c = "this is a new value"
	fmt.Println(b)

	c := myStruct{c: "jifejiofjeoi", b: 83294.234432}

	fmt.Println(c)
}

func array_sliceTest() {
	fmt.Printf("==============Array======================\n")
	// a int array with 10 slots
	var a [10]int
	fmt.Println(a)

	b := []float32{12.152, 1232.156, 92304.99} // I think this is a "slice", whatever that is... okay its a dynamic array
	fmt.Println(b)
	fmt.Printf("The type of b is %T\nThe type of a is %T\n", b, a)

	c := [3]float32{12.152, 1232.156}
	fmt.Println(c)

	fmt.Printf("==============SLICE======================\n")
	d := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(d)

	var myslice []int = d[4:8]
	fmt.Println(myslice)
	fmt.Printf("The type of my slize is %T\n", myslice)

}

func main() {
	variableTests()
	forLoopTests()
	ifTests()
	switchTest()
	pointerTest()
	structTest()
	array_sliceTest()
}
