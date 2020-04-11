package main

import "testing"

func TestFindLeaders(t *testing.T) {
	testData := []struct {
		input    []int
		expected []int
	}{
		{[]int{16, 17, 4, 3, 5, 2}, []int{2, 5, 17}},
		{[]int{1, 2, 3, 4, 0}, []int{0, 4}},
		{[]int{7, 4, 5, 7, 3}, []int{3, 7, 7}},
	}
	for _, data := range testData {
		actual := FindLeaders(data.input)
		if areSlicesEqual(actual, data.expected) {
			t.Log("SUCCESS")
		} else {
			t.Errorf("Fail, input array starts with %d", data.input[0])
		}
	}
}

func areSlicesEqual(actual []int, expected []int) bool {
	if len(actual) != len(expected) {
		return false
	}
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			return false
		}
	}
	return true
}
