package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hello World!")
	// array

	array := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	fmt.Println("=========================")
	arrayA := [5]int{1: 100, 2: 300, 4: 30}
	arrayA[1] += 22
	for i := 0; i < len(arrayA); i++ {
		fmt.Println(arrayA[i])
	}

	fmt.Println("==========pointer===============")
	arrayB := [5]*int{1: new(int), 2: new(int)}
	*arrayB[1] = 10
	*arrayB[2] = 20
	for i := 0; i < len(arrayB); i++ {
		if arrayB[i] != nil {
			fmt.Println(*arrayB[i])
		}
	}

	fmt.Println("============================")
	arrayC := [...]int{1: 100, 2: 300, 4: 30}
	fmt.Println(len(arrayC))

	fmt.Println("============================")
	var arrayD [3]*string
	arrayE := [3]*string{new(string), new(string), new(string)}
	*arrayE[0] = "Happy"
	*arrayE[1] = "New"
	*arrayE[2] = "Year"
	arrayD = arrayE
	for i := 0; i < len(arrayD); i++ {
		fmt.Println(*arrayD[i])
	}

	fmt.Println("=============Multi-dimension-array================")
	//var arrayF [4][2]int
	arrayF := [4][2]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	for i := 0; i < len(arrayF); i++ {
		fmt.Println(arrayF[i][0])
	}

	fmt.Println("==============================")
	arrayG := [3]int{1, 2, 3}
	foo(&arrayG)

	fmt.Println("*************Slice**************")
	slice := make([]int, 3, 5) // slice
	sliceA := []int{1, 2, 3}   // slice
	arrayH := [3]int{1, 2, 3}  // array
	sliceB := []string{99: ""} // slice
	var sliceC []int           // slice: nil
	sliceD := make([]int, 0)   // slice: empty
	fmt.Println(cap(slice), cap(sliceA), cap(arrayH), cap(sliceB), cap(sliceC), cap(sliceD))

	fmt.Println("================================")
	sliceE := []int{1, 2, 3, 4, 5}
	sliceF := sliceE[1:3]
	fmt.Println("sliceF[0]:", sliceF[0], " cap:", cap(sliceF), " len:", len(sliceF))
	sliceF[0] = 100
	fmt.Println(sliceE[1]) // two slices share a same array, it will change when elements of array change
	sliceF = append(sliceF, 20)
	fmt.Println(" cap:", cap(sliceF), " len:", len(sliceF))
	fmt.Println("sliceF[2]:", sliceF[2], "sliceE[3]:", sliceE[3]) // if array have enough capacity , it will use same sub array
	fmt.Println("sliceF[2]:", &sliceF[2], "sliceE[3]:", &sliceE[3])
	sliceF = append(sliceF, 30)
	sliceF = append(sliceF, 40)
	sliceF = append(sliceF, 50)
	sliceF = append(sliceF, 60)
	fmt.Println("sliceF[2]:", &sliceF[2], "sliceE[3]:", &sliceE[3]) // if sub array don't have enough capacity , sliceF will point a new array with copied values.

	fmt.Println("===============================")
	sliceG := sliceF[2:3:4]
	fmt.Println("Len: ", len(sliceG), "Cap:", cap(sliceG)) // Best Practice: cap = len, slice will point to new array when append new elements into slice

	fmt.Println("===============================")
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	fmt.Printf("%v\n", append(s1, s2...)) // ... is an operator, used to add elements to another slice from a slice

	fmt.Println("===============================")
	for index, value := range s1 {
		fmt.Printf("Index: %d, value: %d\n", index, value)
		fmt.Println("original:", &s1[index], " range:", &value) // range operator will create a copy for elements of slice
	}

	fmt.Println("===============================")
	sliceH := [][]int{{10}, {100, 200}}
	sliceH[0] = append(sliceH[0], 20)
	for index, value := range sliceH[0] {
		fmt.Printf("Index: %d, value: %d\n", index, value)
	}
	foo2(s1) //pass slice to func

	fmt.Println("*************MAP**************")
	dict01 := make(map[string]int)
	dict01["A"] = 0
	dict02 := map[string]string{"A": "Java", "B": "Go", "C": "Python"}
	fmt.Println(dict02["A"])
	fmt.Println(dict01["A"])
	value, exists := dict01["B"]
	if exists {
		fmt.Println(value)
	} else {
		fmt.Println("key does not exist")
	}

	//dict03 := map[[]string]int{} // slice is an invalid map key
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}
	for key, value := range colors {
		fmt.Println(key, value)
	}
	removeColor(colors, "red")
	fmt.Println("-------------------")
	for key, value := range colors {
		fmt.Println(key, value)
	}

}

func foo(arr *[3]int) {
	for i := 0; i < 3; i++ {
		fmt.Println(arr[i])
	}
}

func foo2(slice []int) {
	for _, value := range slice {
		fmt.Println(value)
	}
}

func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}
