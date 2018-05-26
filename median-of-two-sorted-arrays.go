package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	total := m + n

	if total%2 == 1 {
		return float64(findKth(nums1, m, nums2, n, total/2+1))
	} else {
		return (float64(findKth(nums1, m, nums2, n,
			total/2)) + float64(findKth(nums1, m, nums2, n, total/2+1))) / 2
	}
}

func findKth(a []int, m int, b []int, n int, k int) int {
	// always assume that m is equal or smaller than n
	if m > n {
		return findKth(b, n, a, m, k)
	}
	if m == 0 {
		return b[k-1]
	}
	if k == 1 {
		if a[0] < b[0] {
			return a[0]
		} else {
			return b[0]
		}
	}

	var pa int
	half := k / 2
	if half < m {
		pa = half
	} else {
		pa = m
	}

	pb := k - pa
	if a[pa-1] < b[pb-1] {
		return findKth(a[pa:], len(a[pa:]), b, n, k-pa)
	} else if a[pa-1] > b[pb-1] {
		return findKth(a, m, b[pb:], len(b[pb:]), k-pb)
	} else {
		return a[pa-1]
	}
}

func main() {
	fmt.Printf("res=%f\n", findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Printf("res=%f\n", findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
