package mysort

import (
	"algorithm/helper"
	"testing"

	"github.com/thoas/go-funk"
)

func TestMedianFinder(t *testing.T) {
	tests := []struct {
		input  []string
		params []int
		want   []float64
	}{
		{
			[]string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"},
			[]int{helper.Invalid, 1, 2, helper.Invalid, 3, helper.Invalid},
			[]float64{helper.InvalidFloat64, helper.InvalidFloat64, helper.InvalidFloat64, 1.5, helper.InvalidFloat64, 2.0},
		},
		{
			[]string{"MedianFinder", "addNum", "findMedian", "addNum", "findMedian"},
			[]int{helper.Invalid, 2, helper.Invalid, 3, helper.Invalid},
			[]float64{helper.InvalidFloat64, helper.InvalidFloat64, 2.0, helper.InvalidFloat64, 2.5},
		},
	}

	for _, test := range tests {
		var m MedianFinder
		got := []float64{}
		for index, operator := range test.input {
			switch operator {
			case "MedianFinder":
				m = Constructor()
				got = append(got, helper.InvalidFloat64)
			case "addNum":
				m.AddNum(test.params[index])
				got = append(got, helper.InvalidFloat64)
			case "findMedian":
				got = append(got, m.FindMedian())
			}
		}

		if !funk.Equal(got, test.want) {
			t.Errorf("test struct MedianFinder failed, input:%v, params:%v, want:%v,got:%v",
				test.input, test.params, test.want, got)
		}
	}
}
