package main

import "fmt"

func main() {
	//arr := []int{16, 17, 4, 3, 5, 2}
	arr := []int{1, 2, 3, 4, 0}
	leaders := FindLeaders(arr)
	for i := len(leaders) - 1; i > 0; i-- {
		fmt.Printf("%d", leaders[i])
		if i > 0 {
			fmt.Print(" ")
		}
	}
}

//FindLeaders test
func FindLeaders(arr []int) []int {
	leaders := make([]int, 1)
	for j := len(arr) - 1; j >= 0; j-- {
		if j == len(arr)-1 {
			leaders[0] = arr[j]
			continue
		}

		if arr[j] >= leaders[(len(leaders)-1)] {
			leaders = append(leaders, arr[j])
		}
	}
	return leaders
}
