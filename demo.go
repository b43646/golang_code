package gotest

import (
"fmt"
_ "github.com/goinaction/code/chapter3/dbdriver/postgres"
)

func main() {
	fmt.Println(123)
	var slice []int
	slice = append(slice, 1)
	slice = append(slice, 2)
	fmt.Println(cap(slice))
	sa := []int{1,2,3,4,5}
	newa := sa[1:3]
	fmt.Println(cap(newa))
}
