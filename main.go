package main

import (
	"fmt"
)

/*
 * Complete the 'arrayManipulation' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. 2D_INTEGER_ARRAY queries
 */

func arrayManipulation(n int32, queries [][]int32) int64 {

	arr := make([]int32, n+1)
	fmt.Println(arr)

	// Apply each operation in an optimized way
	for _, query := range queries {
		a, b, k := query[0]-1, query[1], query[2]

		// Add the value k to each element between indices a and b, inclusive
		arr[a] += k
		if b < n { // Ensure we don't go out of bounds
			arr[b] -= k
		}

	}

	maxVal := int64(0)
	currentSum := int64(0)
	for i := int32(0); i < n; i++ {
		currentSum += int64(arr[i])
		if currentSum > maxVal {
			maxVal = currentSum
		}
	}

	return maxVal

}

func main() {

	n := int32(10)
	queries := [][]int32{
		{1, 5, 3},
		{4, 8, 7},
		{6, 9, 1},
	}

	result := arrayManipulation(n, queries)

	fmt.Println(result)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
