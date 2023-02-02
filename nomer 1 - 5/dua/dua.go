package main

import (
	"fmt"
)

func OneEditApart(a, b string) bool {
	if len(a) == len(b) {
		return oneReplace(a, b)
	} else if len(a)+1 == len(b) {
		return oneInsert(a, b)
	} else if len(a)-1 == len(b) {
		return oneInsert(b, a)
	}
	return false
}
func oneReplace(a, b string) bool {
	foundDifference := false
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			if foundDifference {
				return false
			}
			foundDifference = true
		}
	}
	return true
}
func oneInsert(a, b string) bool {
	index1 := 0
	index2 := 0
	for index2 < len(b) && index1 < len(a) {
		if a[index1] != b[index2] {
			if index1 != index2 {
				return false
			}
			index2++
		} else {
			index1++
			index2++
		}
	}
	return true
}
func main() {
	fmt.Println(OneEditApart("telkom", "telecom"))
	fmt.Println(OneEditApart("telkom", "tlkom"))
}
