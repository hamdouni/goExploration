// Package sum : utilities to sum integers
package sum

// Sum returns sum of integers list
func Sum(v ...int) int {
	if len(v) == 0 {
		return 0
	}
	return v[0] + Sum(v[1:]...)
}
