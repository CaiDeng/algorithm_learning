package helper

import (
	"testing"

	"github.com/thoas/go-funk"
)

func TestNewHeap(t *testing.T) {
	tests := []struct {
		input    []int
		compare  func(int, int) bool
		operates []string
		want     []int
	}{
		{
			[]int{1, 6, 3, 2, 5, 4},
			func(i1, i2 int) bool {
				return i1 > i2
			},
			[]string{"new", "get", "delete", "get", "delete"},
			[]int{invalid, 6, 6, 5, 5},
		},
		{
			[]int{1, 6, 3, 2, 5, 4},
			func(i1, i2 int) bool {
				return i1 < i2
			},
			[]string{"new", "get", "delete", "delete", "get"},
			[]int{invalid, 1, 1, 2, 3},
		},
	}

Loop:
	for _, test := range tests {
		var h Heap
		got := []int{}
		for _, opreate := range test.operates {
			switch opreate {
			case "new":
				var err error
				h, err = NewHeap(test.input, test.compare)
				if err != nil {
					break Loop
				}
				got = append(got, invalid)
			case "get":
				if extremum, err := h.Extremum(); err != nil {
					got = append(got, invalid)
				} else {
					got = append(got, extremum)
				}
			case "delete":
				if extremum, err := h.DeleteExtremum(); err != nil {
					got = append(got, invalid)
				} else {
					got = append(got, extremum)
				}
			}
		}

		if !funk.Equal(got, test.want) {
			t.Errorf("test function heap's new ,get and delete failed, input:%v,operator:%v,want:%v, got:%v",
				test.input, test.operates, test.want, got)
		}
	}

}

func TestSet(t *testing.T) {
	tests := []struct {
		input   []int
		compare func(int, int) bool
		params  [][2]int
		want    []int
	}{
		{
			[]int{6, 2, 4, 1, 3},
			func(i1, i2 int) bool {
				return i1 > i2
			},
			[][2]int{
				{0, -1},
				{0, 9},
				{1, -2},
				{1, 10},
				{2, -3},
				{2, 11},
			},
			[]int{11, 9, 10, 1, -2},
		},
		{
			[]int{5, 2, 4, 1, 7},
			func(i1, i2 int) bool {
				return i1 < i2
			},
			[][2]int{
				{0, 9},
				{0, -1},
				{1, 10},
				{1, -2},
				{2, 11},
				{2, -3},
			},
			[]int{-3, -1, -2, 9, 10},
		},
	}

	for _, test := range tests {
		h, err := NewHeap(test.input, test.compare)
		if err != nil {
			t.Errorf("test function Set failed, reason:%v", err)
		}

		for _, param := range test.params {
			h.Set(param[0], param[1])
		}

		got := h.elem
		if !funk.Equal(got, test.want) {
			t.Errorf("test function Set failed, input:%v,params:%v,want:%v, got:%v",
				test.input, test.params, test.want, got)
		}
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		input   []int
		compare func(int, int) bool
		params  []int
		want    []int
	}{
		{
			[]int{6, 2, 4, 1, 3},
			func(i1, i2 int) bool {
				return i1 > i2
			},
			[]int{
				9,
				-1,
				10,
				-2,
			},
			[]int{10, 9, 6, 3, 2, 4, -1, 1, -2},
		},
		{
			[]int{5, 2, 4, 1, 7},
			func(i1, i2 int) bool {
				return i1 < i2
			},
			[]int{
				-1,
				9,
				-2,
				10,
			},
			[]int{-2, -1, 1, 2, 7, 4, 9, 5, 10},
		},
	}

	for _, test := range tests {
		h, err := NewHeap(test.input, test.compare)
		if err != nil {
			t.Errorf("test function Set failed, reason:%v", err)
		}

		for _, param := range test.params {
			h.Insert(param)
		}

		got := h.elem
		if !funk.Equal(got, test.want) {
			t.Errorf("test function Set failed, input:%v,params:%v,want:%v, got:%v",
				test.input, test.params, test.want, got)
		}
	}
}

func TestSort(t *testing.T) {
	tests := []struct {
		input   []int
		compare func(int, int) bool
		want    []int
	}{
		{
			[]int{6, 2, 4, 1, 3, 0, 9},
			func(i1, i2 int) bool {
				return i1 > i2
			},
			[]int{0, 1, 2, 3, 4, 6, 9},
		},
		{
			[]int{6, 2, 4, 1, 3},
			func(i1, i2 int) bool {
				return i1 < i2
			},
			[]int{6, 4, 3, 2, 1},
		},
	}

	for _, test := range tests {
		h, err := NewHeap(test.input, test.compare)
		if err != nil {
			t.Errorf("test function Sort failed, reason:%v", err)
		}

		got, err := h.Sort()
		if err != nil {
			t.Errorf("test function Sort failed, reason:%v", err)
		}

		if !funk.Equal(got, test.want) {
			t.Errorf("test function Set failed, input:%v,want:%v, got:%v",
				test.input, test.want, got)
		}
	}
}
