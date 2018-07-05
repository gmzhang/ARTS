package main

import (
	"errors"
	"fmt"
)

/**
There are two sorted arrays nums1 and nums2 of size m and n respectively.
Find the median of the two sorted arrays. The overall run time complexity
should be O(log(m+n)).

Example 1:

nums1 = [1, 3]
nums2 = [2]
The median is 2.0

Example 2:

nums1 = [1, 2]
nums2 = [3, 4]
The median is (2 + 3)/2 = 2.5


归并排序法
如果没有时间约束，就可以使用归并排序，若排序后的结果为nums3,元素个数为k=m+n;最终将问题转化为：
求数组nums3的中位数。

若k为奇数，直接为nums3[k/2];
若k为偶数，则为(nums3[k/2]+nums3[k/2-1])/2；

 */
func main() {
	nums1 := []float32{1, 2}
	nums2 := []float32{3, 4}
	ret, _ := FindMedianSortedArrays(nums1, nums2)
	fmt.Println(ret)

}

func findMedianSortedSignleArray(nums []float32) (result float32, err error) {
	if nums == nil || len(nums) == 0 {
		return 0, errors.New("sorted array are illegal")
	}

	len := len(nums)

	if (len % 2) == 0 {
		return (nums[len/2] + nums[len/2-1]) / 2, nil
	}

	return nums[len/2], nil
}

func FindMedianSortedArrays(nums1 []float32, nums2 []float32) (result float32, err error) {
	if nums1 == nil || len(nums1) == 0 {
		return findMedianSortedSignleArray(nums2)
	}

	if nums2 == nil || len(nums2) == 0 {
		return findMedianSortedSignleArray(nums1)
	}

	n := len(nums1)
	m := len(nums2)

	newNums := make([]float32, 0)

	i := 0
	j := 0

	for (i < n && j < m) {
		if nums1[i] > nums2[j] {
			newNums = append(newNums, nums2[j])
			j = j + 1
		} else {
			newNums = append(newNums, nums1[i])
			i = i + 1
		}
	}

	for i < n {
		newNums = append(newNums, nums1[i])
		i = i + 1
	}

	for j < m {
		newNums = append(newNums, nums2[j])
		j = j + 1
	}

	return findMedianSortedSignleArray(newNums)

}
