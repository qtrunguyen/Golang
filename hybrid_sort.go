package main

import (
	"math/rand"
)

/* A hybridsort between Insertion and QuickSort
Quicksort performs well with large data, while Insertion Sort works better with small data
*/
const insertionSortThreshhold = 10

func HybridSort(arr []int) {
	QuickSort(arr, 0, len(arr)-1)
}

func Partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func QuickSort(arr []int, low int, high int) {
	if low < high {
		if high-low+1 <= insertionSortThreshhold {
			InsertionSort(arr, low, high)
		} else {
			pivotIndex := Partition(arr, low, high)
			QuickSort(arr, low, pivotIndex-1)
			QuickSort(arr, pivotIndex+1, high)
		}
	}
}

func InsertionSort(arr []int, low int, high int) {
	for i := low + 1; i <= high; i++ {
		key := arr[i]
		j := i - 1
		for j >= low && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func GenerateRandomArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(10000)
	}
	return arr
}

func main() {
	rand.Seed(time.Now().UnixNano())
	arr := GenerateRandomArray(100)
	fmt.Println("Original array:", arr)
	HybridSort(arr)
	fmt.Println("Sorted array:", arr)
}
