package main

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	//size := m + n
	//ans := make([]int, size)
	a := nums1[:m]
	b := nums2[:n]
	sort.Ints(a)
	sort.Ints(b)

	a = append(a, b...)

	sort.Ints(a)
}
