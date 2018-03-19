package main

import (
	"fmt"
	"math/rand"
	"math"
)

func main(){
	fmt.Println("hello, world")
	fmt.Println("My favorite no is : ", rand.Int())
	fmt.Println(math.Pi)
	fmt.Println(add(2,3))

	a,b := swap("hello",	"world")
	fmt.Println(a,b)

	var c, d int
	c, d = split(12)
	fmt.Println(c,d)

	var e,f int = 10, 20
	var g, h, i = false, true, "hi"
	fmt.Println(e, f, g, h, i)

	Sqrt(2)

}

func add(x, y int) int{
	return x+y
}

func swap(a, b string) (string, string) {
	return b, a
}

func split(x int)	(a, b int)  {
	a = x -2;
	b = x - a
	return
}

func Sqrt(x float64)  {
	var z float64 = 1
	i :=1
	for {
		prevZ := z;
		z -= (z*z - x)/2*z
		fmt.Println("Current Z is ", z)
		if math.Mod(prevZ, z) < .1 && i>1 {
			break
		}
		i++
	}
	
}