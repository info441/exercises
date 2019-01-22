package mergesort

import (
	"reflect"
	"testing"
)

// TODO: Add example of cases below and test them
func TestMergeSort(t *testing.T) {
	cases := []struct {
		input          []int
		expectedOutput []int
	}{
		{
			// Add example of cases here
		},
	}
	
	for _, c := range cases {
		output := MergeSort(c.input)
		if !reflect.DeepEqual(output, c.expectedOutput) {
			t.Errorf("incorrect output for %v: expected %v but got %v", c.input, c.expectedOutput, output)
		}
	}
}