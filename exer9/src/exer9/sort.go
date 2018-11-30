package exer9

import (
	"math/rand"
)

// Partition the slice arr around a random pivot (in-place), and return the pivot location.
func partition(arr []float64) int {
	// Adapted from https://stackoverflow.com/a/15803401/6871666
	left := 0
	right := len(arr) - 1

	// Choose random pivot
	pivotIndex := rand.Intn(len(arr))

	// Stash pivot at the right of the slice
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Move elements smaller than the pivot to the left
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Place the pivot after the last-smaller element
	arr[left], arr[right] = arr[right], arr[left]
	return left
}

func InsertionSort(arr []float64) {
	length := len(arr)
	for i:=1; i<length; i++{
		key := arr[i]
		j := i-1
		for j>=0 && key < arr[j]{
			arr[j+1] = arr[j]
			j = j-1
		}
		arr[j+1] = key
	}
	// TODO: implement insertion sort
}

const insertionSortCutoff = 1000

func Quick(arr []float64){
	low := 0
	high := len(arr) -1
	pi := partition(arr)
	Quick(arr[low:pi])
	Quick(arr[(pi+1):high])
}

func QuickSort(arr []float64) {
	// TODO: implement Quicksort:
	//   do nothing for length < 2
	length := len(arr)
	if length <insertionSortCutoff{
		InsertionSort(arr)
	} else {
		Quick(arr)
	}
	//   do insertion sort for length < insertionSortCutoff
	//   do Quicksort otherwise.
	// TODO: decide on a good value for insertionSortCutoff
}
