package mysort

import (
	"bytes"
	"sort"
	"strconv"
)

func minNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		return strconv.Itoa(nums[i])+strconv.Itoa(nums[j]) < strconv.Itoa(nums[j])+strconv.Itoa(nums[i])
	})

	var buffer bytes.Buffer
	for _, value := range nums {
		buffer.WriteString(strconv.Itoa(value))
	}
	return buffer.String()
}
