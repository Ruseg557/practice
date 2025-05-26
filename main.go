package main

import (
	"fmt"
)

func merge(arr1, arr2 []int) []int {
	var i, j = 0, 0
	var c = []int{}
	for i < len(arr1) || j < len(arr2) {
		if i < len(arr1) && (j >= len(arr2) || arr1[i] < arr2[j]) {
			c = append(c, arr1[i])
			i++
		} else {
			c = append(c, arr2[j])
			j++
		}
	}
	return c
}

func merge_sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	var arr1 = merge_sort(arr[:mid])
	var arr2 = merge_sort(arr[mid:])
	return merge(arr1, arr2)
}

func main() {
	var arr = []int{5, 2, -1, 0, 5, 16, 27, 8, 9, 10}
	var sortedArr = merge_sort(arr)
	fmt.Println(arr)
	fmt.Println(sortedArr)

}
