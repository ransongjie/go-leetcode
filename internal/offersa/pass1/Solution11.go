package pass1
/**
 * 剑指 Offer II 011. 0和1个数相同的连续最长子数组
 * - 含有相同数量的0和1的子数组
 * - 子数组连续
 * - 最长子数组
 */
type Solution11 struct{}
    /**
     * 问题转化，将0看成-1,0和1个数相同的子数组，-1和1个数相同的子数组，sum=0
     * 前缀和：map<当前和,索引>, 前缀和索引,中间和=0,当前和索引
     * 
     * @param nums
     * @return
     */
func (solution11 Solution11) findMaxLength(nums []int) int {
	// 结果
	var maxLen int=0
	// 定义
	var am map[int]int=make(map[int]int)
	var sum int=0
	// 初始化
	am[sum]=-1

	for i,num:=range nums{
		// 求和
		if num==0{
			sum--
		}else{
			sum++
		}
		// 判断前缀和
		preSumIdx,ok:=am[sum-0]
		if ok{
			if maxLen<(i-preSumIdx){
				maxLen=i-preSumIdx
			}
		}else{
			// xcrj else更新, currSum-0=preSum, currSum=preSum, 但索引更大
			am[sum]=i
		}
	}

	return maxLen
}