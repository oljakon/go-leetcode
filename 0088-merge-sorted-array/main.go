package main

func merge1(nums1 []int, m int, nums2 []int, n int) {

	//     i, j, k := m-1, n-1, m+n-1
	//     for ; j >= 0; k-- {
	//         if i >= 0 && nums1[i] > nums2[j] {
	//             nums1[k] = nums1[i]
	//             i--
	//         } else {
	//             nums1[k] = nums2[j]
	//             j--
	//         }
	//     }

	last := m + n - 1
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[last] = nums1[m-1]
			m--
		} else {
			nums1[last] = nums2[n-1]
			n--
		}
		last--
	}

	for n > 0 {
		nums1[last] = nums2[n-1]
		n--
		last--
	}
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for ; j >= 0; k-- {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}
