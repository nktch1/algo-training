package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	left, right, mid := 0, n, 0

	for left <= right {
		mid = (left + right) / 2

		var (
			fM  = isBadVersion(mid)
			fM1 = isBadVersion(mid + 1)
		)

		if !fM && fM1 {
			break
		}

		if fM1 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return mid + 1
}

func isBadVersion(mid int) bool { return false }
