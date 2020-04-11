package main

import "fmt"

func main() {
	//arr := []int{16, 17, 4, 3, 5, 2}
	arr := []int{1, 2, 3, 4, 0}
	leaders := Find(arr)
	for i := 0; i < len(arr); i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%d", leaders[i])
	}
}

//Find test
func Find(arr []int) []int {
	leaders := make([]int, 0)
	arr = findLeaders(arr)
	for i := 0; i < len(arr); i++ {
		if arr[i] != -1 {
			leaders = append(leaders, arr[i])
		}
	}
	return leaders
}

//FindLeaders test
func findLeaders(arr []int) []int {
	//leaders := make([]int, 1)
	max := arr[len(arr)-1]
	for j := len(arr) - 2; j >= 0; j-- {
		if arr[j] >= max {
			max = arr[j]
		} else {
			arr[j] = -1
		}
	}
	return arr
}
