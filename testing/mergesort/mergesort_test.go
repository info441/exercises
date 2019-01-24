package mergesort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	cases := []struct {
		input          []int
		expectedOutput []int
	}{
		{
			[]int{3,2,1},
			[]int{1,2,3},
		},
		{
			[]int{0, 2},
			[]int{0, 2},
		},
		{
			[]int{9,8,7,6,5,4,3,2,1},
			[]int{1,2,3,4,5,6,7,8,9},
		},
	}
	
	for _, c := range cases {
		output := MergeSort(c.input)
		if !reflect.DeepEqual(output, c.expectedOutput) {
			t.Errorf("incorrect output for %v: expected %v but got %v", c.input, c.expectedOutput, output)
		}
	}
}