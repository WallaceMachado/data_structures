package arrays

import (
	"testing"
)

func TestPrintArray(t *testing.T) {
	tests := []struct {
		input  []int
		output string
	}{
		{[]int{1, 2, 3}, "1\t2\t3"},
		{[]int{4, 5, 6}, "4\t5\t6"},
		{[]int{7, 8, 9, 10}, "7\t8\t9\t10"},
		{[]int{}, ""},
	}

	for _, test := range tests {
		result := printArray(test.input)
		if result != test.output {
			t.Errorf("printArray(%v) = %s; want %s", test.input, result, test.output)
		}
	}
}
