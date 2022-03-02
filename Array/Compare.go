/**
 * @Author: luccifer
 * @Description:
 * @File:  MinInTwo
 * @Version: 1.0.0
 * @Date: 2022/3/2 16:39
 */

package Array

func MinOrMax(a, b int, f func(int, int) bool) int {
	if f(a, b) {
		return a
	}
	return b
}

func GetMax(a, b int) bool {
	return a > b
}

func GetMin(a, b int) bool {
	return a < b
}