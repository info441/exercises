package trie

import (
	"reflect"
	"sort"
	"testing"
)

func TestInt64SetAddHas(t *testing.T) {
	cases := []struct {
		name     string
		values   []int64
		expected []int64
	}{
		{
			"Single Value",
			[]int64{1},
			[]int64{1},
		},
		{
			"Duplicate Values",
			[]int64{1, 1},
			[]int64{1},
		},
		{
			"Distinct Values",
			[]int64{1, 2, 3},
			[]int64{1, 2, 3},
		},
		{
			"Distinct Values Then Duplicates",
			[]int64{1, 2, 3, 1, 2, 3},
			[]int64{1, 2, 3},
		},
		{
			"Duplicates Then Distinct Values",
			[]int64{1, 1, 2, 2, 3, 3},
			[]int64{1, 2, 3},
		},
		{
			"Intermixed",
			[]int64{1, 2, 1, 3, 4, 3},
			[]int64{1, 2, 3, 4},
		},
	}
	
	for _, c := range cases {
		testset := int64set{}
		for _, v := range c.values {
			expectedRet := !testset.has(v)
			ret := testset.add(v)
			if expectedRet != ret {
				t.Errorf("case %s: incorrect return value when adding %d: expected %t but got %t",
					c.name, v, expectedRet, ret)
			}
		}
		if len(testset) != len(c.expected) {
			t.Errorf("case %s: incorrect length: expected %d but got %d",
				c.name, len(c.expected), len(testset))
		}
		for _, v := range c.expected {
			if !testset.has(v) {
				t.Errorf("case %s: expected value %d is not in the set",
					c.name, v)
			}
		}
	}
}

func TestInt64SetRemove(t *testing.T) {
	cases := []struct {
		name     string
		values   []int64
		toRemove int64
		expected []int64
	}{
		{
			"One Removed from Many",
			[]int64{1, 2, 3},
			2,
			[]int64{1, 3},
		},
		{
			"Last One",
			[]int64{1},
			1,
			[]int64{},
		},
		{
			"Not Found",
			[]int64{1},
			2,
			[]int64{1},
		},
		{
			"Empty",
			[]int64{},
			2,
			[]int64{},
		},
	}
	
	for _, c := range cases {
		testset := int64set{}
		for _, v := range c.values {
			testset.add(v)
		}
		expectedRet := testset.has(c.toRemove)
		ret := testset.remove(c.toRemove)
		if expectedRet != ret {
			t.Errorf("case %s: incorrect return value when removing %d: expected %t but got %t",
				c.name, c.toRemove, expectedRet, ret)
		}
		if len(testset) != len(c.expected) {
			t.Errorf("case %s: incorrect length after remove: expected %d but got %d",
				c.name, len(c.expected), len(testset))
		}
		for _, v := range c.expected {
			if !testset.has(v) {
				t.Errorf("case %s: expected value %d was not in set after remove",
					c.name, v)
			}
		}
	}
}

func TestInt64SetAll(t *testing.T) {
	cases := []struct {
		name     string
		values   []int64
		expected []int64
	}{
		{
			"Multiple Distinct Entries",
			[]int64{1, 2, 3},
			[]int64{1, 2, 3},
		},
		{
			"Duplicate Entries",
			[]int64{1, 1, 2, 3, 2, 3},
			[]int64{1, 2, 3},
		},
		{
			"Single Entry",
			[]int64{1},
			[]int64{1},
		},
		{
			"Empty Set",
			[]int64{},
			[]int64{},
		},
	}
	
	for _, c := range cases {
		testset := int64set{}
		for _, v := range c.values {
			testset.add(v)
		}
		
		actual := testset.all()
		//since the order is random, sort so that we can compare
		sort.Slice(actual, func(i, j int) bool { return actual[i] < actual[j] })
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("case %s: incorrect results: expected %v but got %v", c.name, c.expected, actual)
		}
	}
}