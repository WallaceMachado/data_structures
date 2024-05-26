package arrays

import "fmt"

func printArray(arr []int) string {
	result := ""
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d\t", i)
		result += fmt.Sprintf("%d", arr[i])
	}
	return result
}
