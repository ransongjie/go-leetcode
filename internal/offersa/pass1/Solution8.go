package pass1
import (
	"math"
)
/**
 * 剑指 Offer II 008. 和大于等于target的连续最短子数组
 * - 元素和>=target
 * - 元素构成的子序列是连续的
 * - 子序列的长度最短
 */

type Solution8 struct{}

    /**
     * 双指针同向移动
     * 两个while循环夹击出最短子数组
     * - 先j从0开始往右移动，以满足sum>=target
     * - 在i从0开始往右移动，以满足len最短的 sum>=target
     * @param target
     * @param nums
     * @return
     */
func (solution8 Solution8) minSubArrayLen(target int, nums []int) int {
	var i int=0
	var j int=0
	var sum int=0
	var r int=math.MaxInt
	for j<len(nums){
		sum+=nums[j]
		for sum>=target{
			if r>(j-i+1){
				r=(j-i+1)
			}
			sum-=nums[i]
			i++
		}

		j++
	}

	if r==math.MaxInt {
		return 0
	}else{
		return r
	}
}