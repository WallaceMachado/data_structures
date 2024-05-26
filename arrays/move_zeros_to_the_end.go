package arrays

func moveZeros(arr []int) []int {
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 && arr[j] == 0 {
			arr[i], arr[j] = arr[j], arr[i]
		}

		if arr[j] != 0 {
			j++
		}
	}

	return arr
}
