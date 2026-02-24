package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {
	fmt.Println("Hello, World!")

	fmt.Println("Num CPU", runtime.NumCPU())

	// variables types
	var num1 int8 = 20
	num2 := 1000
	var num3 float32 = 3.14

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Printf("%T", num2)
	fmt.Println(reflect.TypeOf(num2))

	fmt.Println(num3)

	var str1 string = "abc"
	fmt.Println(str1)

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	i := 1
	for i < 5 {
		fmt.Println(i)
		i++
	}
	fmt.Println(isAdult(19))
	fmt.Println(isAdult(18))

	// Looping over slices
	// nums := []int{10, 20, 30}
	var nums []int = []int{10, 20, 30}
	for index, value := range nums {
		fmt.Println(index, value)
	}

	ages := map[string]int8{
		"alice": 20,
		"john":  43,
		"jane":  50,
	}

	fmt.Println("Traversing through the map ages")
	for index, value := range ages {
		fmt.Println(index, value)
	}
}

func isAdult(age int) bool {
	if age > 18 {
		return true
	} else {
		return false
	}

}
