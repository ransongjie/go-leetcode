package pass1

/**
 * 剑指 Offer II 009. 乘积小于k的最短连续子数组的数量
 * - 数组元素乘积小于k
 * - 数组元素连续
 * - 有多少个这样的子数组
 */

type Solution9 struct{}

    /**
     * 乘积小于k的最短连续子数组的数量
     * 两个while循环夹击
     * - 先j从0到len 探求满足 p>=k 的解
     * - 再i从0到j   探求满足 最短 p<k 的解
     * @param nums
     * @param k
     * @return
     */
func (solution9 Solution9) numSubarrayProductLessThanK(nums []int, k int) int {
	var r int=0
	var i int=0
	var j int=0
	//乘法从1开始
	var p int=1

	for j<len(nums) {
		p*=nums[j]
		for i<=j && p>=k{
			p/=nums[i]
			i++
		}
		// 满足 p<k 的连续子数组,连续子数组的子数组也满足 p<k
		r+=j-i+1
		j++
	}

	return r
}