package main

import (
	"fmt"
	"log"
	"math/rand"
)

// Nodes a generic type for slice of any type.

//type Nodes[T toWork] []T

// Sort sorts does not mind about the fields or values of T.
func Sort(input []toWork) {
	length := len(input)
	log.Println("length", length)

	var i1, i2 int
	rightGap = float64(length - 1)
	for gap := length - 1; gap > 0; gap -= reducer(gap) {
		for i1, i2 = 0, gap; i2 < length; i1++ {
			if input[i1].id > input[i2].id {
				//log.Println("changing", input[i1], input[i2])
				input[i1], input[i2] = input[i2], input[i1]
			}
			i2++
		}
		if gap > 2000 {
			log.Println("gap", gap)
			log.Println("gap must be", rightGap)
		}
	}
}

var rightGap float64

func reducer(gap int) int {
	rightGap = float64(int((rightGap / 1.3) + .5))
	if gap < 9 {
		return 1
	}
	return gap / 4
}

type toWork struct {
	id int
}

func main() {
	var testSlice = []toWork{
		{id: 4},
		{id: 10},
		{id: 12},
		{id: 2},
		{id: 3},
		{id: 20},
		{id: 22},
		{id: 25},
		{id: 5},
		{id: 9},
		{id: 15},
		{id: 222},
		{id: 1},
	}

	Sort(testSlice)

	log.Println(testSlice)

	var testSlice2 = make([]toWork, 5000)
	for key := range testSlice2 {
		testSlice2[key].id = rand.Int()
	}
	Sort(testSlice2)
	for key := range testSlice2 {
		if key != 0 {
			if testSlice2[key].id < testSlice2[key-1].id {
				fmt.Println("algorithm does not work")
			}
		}
	}
}
