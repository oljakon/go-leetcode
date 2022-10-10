package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool {
	//return false
	return true
}

func firstBadVersion(n int) int {
	firstBad := n

	l, r := 1, n

	for l <= r {
		mid := l + (r-l)/2
		if isBadVersion(mid) {
			firstBad = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return firstBad
}
