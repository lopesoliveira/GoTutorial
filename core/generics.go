package main

func sumSliceIndex(slice []int) int {
	var sum int
	for v := range slice {
		sum += v
	}
	return sum //because for only has one variable it assign it to the index, meaning it sums 0+1+2 = 3
}

func sumSlice(slice []int) int {
	var sum int
	for _, v := range slice {
		sum += v
	}
	return sum //in this for loop v is the value of the slices meaning it will sum 1+2+3 = 6
}

func sumSliceGeneric[T int | float32 | float64 | string](slice []T) T {
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}

// There's also any type. It can't be used here because it does not make sense for booleans

//Check if a slice is empty. This works for whatever type

func isSliceEmpty[T any](slice []T) bool {
	/*if slice == nil {
		return true
	} else {
		return false
	}*/
	return len(slice) == 0
}
