package pass1
/**
 * 剑指 Offer II 010. 和为k的连续子数组的数量
 * - 子数组的和等于k
 * - 子数组连续
 * - 求和为k的连续子数组的个数。
 */
type Solution10 struct{}
    /**
     * 边界指针倒序求解
     * - j做i的右边界
     * - i倒序从j到0，防止重复
     * - xcrj 防止重复 i每次都从1个新的起点j出发，如果i从0一定会重复
     * 
     * @param nums
     * @param k
     * @return
     */
func (solution10 Solution10) subarraySum1(nums []int, k int) int {
	var r=0
	for j:=0;j<len(nums);j+=1{
		var sum=0
		for i:=j;i>=0;i-=1{
			sum+=nums[i]
			if sum==k{
				r++
			}
		}
	}
	
	return r
}
    /**
     * map<sum,次数>
     * 在存储当前和出现次数的过程中寻找前缀和
     * 前缀和+散列表
     * - 当前和-k=前缀和
     * - map<当前和,次数>
     * 
     * 
     * @param nums
     * @param k
     * @return
     */
func (solution10 Solution10) subarraySum2(nums []int, k int) int {
	var r int=0
	var mp map[int]int=make(map[int]int)
	var sum int=0
	mp[sum]=1
	for _,num:=range nums{
		sum+=num
		// contains
		_,ok:=mp[sum-k]
		if ok{
			r+=mp[sum-k]
		}

		// xcrj go get or default
		mp[sum]=mp[sum]+1
	}
	return r
}