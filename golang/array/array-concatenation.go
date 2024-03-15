// https://leetcode.com/problems/concatenation-of-array/
// 9ms
// 6.58MB

package main

func main() {

}

func getConcatenation(nums []int) []int {
	var result []int
	for i := 0; i < 2; i++ {
		for j := 0; j < len(nums); j++ {
			result = append(result, nums[j])
		}
	}
	return result
}
