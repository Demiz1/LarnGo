package main

import (
	"fmt"
	"time"
)
func simpleAdd(x int, y int) int{
	return x+y
}
func twoReturn(x string) (string, string){
	return x,x
}

func main(){
	fmt.Println("Hello world!", time.Now())
	a := 1
	b := 1336
	fmt.Println("now i add two values", simpleAdd(a,b))
	s := "hello"
	var ost string
	_ , ost = twoReturn(s)
	fmt.Println("here i use a function with two returns", ost)

	v := 42 // change me!
	fmt.Printf("v is of type %T\n", uint16(v))
	v_2 := float64(v)
	fmt.Printf("%T\n",v_2)

	var g,j,c float64 = 1,24,5

	fmt.Printf("%T,%T,%T\n",g,j,c)

}