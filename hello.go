package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {

	b := 5

	a := []int{5, b, 3, 2, 1}
	a = append(a, 13)

	fmt.Println(a)

	vertices := make(map[string]int)

	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 12

	delete(vertices, "square")

	fmt.Println(vertices)

	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}

	for i := 0; i < 8; i++ {
		fmt.Println(i)
	}

	result := sum(2, 3)
	fmt.Println(result)

	//this is the struct defined as person above
	p := person{name: "jamke", age: 23}
	fmt.Println(p)
	fmt.Println(p.age)
}

func sum(x int, y int) int {
	return x + y
}
