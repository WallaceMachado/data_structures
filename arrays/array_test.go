package arrays

import (
	"reflect"
	"testing"
)

func TestPrintArray(t *testing.T) {
	tests := []struct {
		input  []int
		output string
	}{
		{[]int{1, 2, 3}, "123"},
		{[]int{4, 5, 6}, "456"},
		{[]int{7, 8, 9, 10}, "78910"},
		{[]int{}, ""},
	}

	for _, test := range tests {
		result := printArray(test.input)
		if result != test.output {
			t.Errorf("printArray(%v) = %s; want %s", test.input, result, test.output)
		}
	}
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestRemoveEvenNumbersFromArray(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 3, 5}},
		{[]int{2, 4, 6, 8}, []int{}},
		{[]int{1, 3, 5, 7}, []int{1, 3, 5, 7}},
		{[]int{}, []int{}},
		{[]int{0, 1, 2, 3, 4}, []int{1, 3}},
	}

	for _, test := range tests {
		result := removeEvenNumbersFromArray(test.input)
		if !slicesEqual(result, test.expected) {
			t.Errorf("removeEvenNumbersFromArray(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestMoveZeros(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0, 0, 1}, []int{1, 0, 0}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{}, []int{}},
		{[]int{1, 0, 1, 0, 1}, []int{1, 1, 1, 0, 0}},
	}

	for _, test := range tests {
		result := moveZeros(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("moveZeros(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestReverseArray(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{2, 4, 6, 8}, []int{8, 6, 4, 2}},
	}

	for _, test := range tests {
		result := reverseArray(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("reverseArray(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestFindMissingNumber(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{0, 1, 2, 4, 5}, 3}, // Number missing in the middle
		{[]int{1, 2, 3, 4, 5}, 0}, // Number missing at the start
		{[]int{0, 1, 3, 4, 5}, 2}, // Number missing in the middle
		{[]int{0, 1, 2, 3, 4}, 5}, // Number missing at the end
		{[]int{1, 0}, 2},          // Small array
		{[]int{2, 3, 1, 0, 5}, 4}, // Unordered array
	}

	for _, test := range tests {
		result := findMissingNumber(test.input)
		if result != test.expected {
			t.Errorf("findMissingNumber(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},            // Palíndromo: string vazia
		{"a", true},           // Palíndromo: uma letra
		{"radar", true},       // Palíndromo: palavra ímpar
		{"level", true},       // Palíndromo: palavra par
		{"hello", false},      // Não palíndromo: ímpar
		{"world", false},      // Não palíndromo: par
		{"abba", true},        // Palíndromo: palavra com repetições
		{"racecar", true},     // Palíndromo: palavra com caracteres diferentes
		{"madam", true},       // Palíndromo: palavra simples
		{"palindrome", false}, // Não palíndromo: palavra simples

	}

	for _, test := range tests {
		result := isPalindrome(test.input)
		if result != test.expected {
			t.Errorf("isPalindrome(%s) = %v; want %v", test.input, result, test.expected)
		}
	}
}
