/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// time: O(log2(n)), space: O(1)
// binary search
func firstBadVersion(n int) int {
    left := 1
    right := n
    for left < right {
        cur := (right - left) / 2 + left
        if isBadVersion(cur) {
            right = cur
        } else {
            left = cur + 1
        }
    }
    return left
}
