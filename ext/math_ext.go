package ext

import (
	"golang.org/x/exp/constraints"
)

func MinByQsort[T constraints.Ordered](args ...T) T {
	qsort(args)
	return args[0]
}
func MaxByQsort[T constraints.Ordered](args ...T) T {
	qsort(args)
	return args[len(args)-1]
}
func Min[T constraints.Ordered](args ...T) T {
	res := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < res {
			res = args[i]
		}
	}
	return res
}
func Max[T constraints.Ordered](args ...T) T {
	res := args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > res {
			res = args[i]
		}
	}
	return res
}

// qsort quick sort
func qsort[T constraints.Ordered](nums []T) {
	if len(nums) <= 1 {
		return
	}
	pivot := nums[0]
	left, right := 0, len(nums)-1
	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= pivot {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	qsort(nums[:left])
	qsort(nums[left+1:])
}
