package main

import "fmt"
// https://play.golang.org/p/LCv9YRFgbVN
// https://leetcode.com/problems/subarray-sum-equals-k/
func main() {
	nums := []int{1}
	k := 1
	x := subarraySum(nums, k)
	fmt.Println("result", x)
}

func subarraySum(nums []int, k int) int {
    var preSumMap = make(map[int]int)
    preSumMap[0] = 1
    var result = 0
    var sum = 0
    for _, value := range nums {
        sum += value
		
        tmp := sum-k
        if _, ok := preSumMap[tmp]; ok {
            result +=  preSumMap[tmp]
        }
        
        if _, ok := preSumMap[sum]; ok {
            preSumMap[sum] += 1
        } else {
		   
            preSumMap[sum] = 1
        }
    }
    return result
}
