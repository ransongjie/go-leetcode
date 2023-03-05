package pass1

/**
 * 二分查找
 * 双指针
 * 剑指 Offer II 006. 从已升序排序数组中寻找两个数，他们的和为target
 * - 数组中只有一对符合条件的数字。
 * - 同一个数字不能使用两次，已经找遍历过的数字不再使用
 */

type Solution6 struct{}

    /**
     * numbers已经有序，使用二分查找
     * - target-numbers[i]是想要寻找的值
     * - 已经找过的值不再找
     * 
     * @param numbers
     * @param target
     * @return new int[0] 空slice
     */
func (solution6 Solution6)  twoSum1(numbers []int ,target int)[]int{
	//遍历数组寻找 target-numbers[i]，已遍历过的数字不再使用
	for i,num:=range numbers{
		var j int=solution6.binarySearch(numbers,target-num,i+1,len(numbers)-1)
		if j!=-1{
			return []int{i,j}
		}
	}
	// xcrj go创建空slice
	return make([]int,0)
}

    /**
     * 二分查找
     * - 寻找值等于中间值返回元素索引
     * - 寻找值大于中间值往右侧寻找
     * - 寻找值小于中间值往左侧寻找
     * 
     * @param numbers
     * @param x
     * @param start
     * @param end
     * @return -1表示没有找到
     */
func (solution6 Solution6) binarySearch(numbers []int,x int,start int, end int)int{
	var i int=start
	var j int=end

	for i<=j{
		var mid int=(i+j)/2
		if numbers[mid]==x{
			return mid
		}
		// xcrj i, x_idx, mid, j
		// xcrj i, x_idx, j, mid
		if numbers[mid]>x{
			j=mid-1
		}else{
			i=mid+1
		}
	}

	return -1
}

    /**
     * 双指针，反向指针
     * - 已知是升序序列
     * - i指针从左往右，j指针从右往左
     * - 元素i+元素j<target，i往右移动
     * - 元素i+元素j>target，j往左移动
     * @param numbers
     * @param target
     * @return new int[0]空数组
     */
func (solution6 Solution6)  twoSum2(numbers []int ,target int)[]int{
	// xcrj 已经遍历过的不能再使用&&numbers已经有序，因此可以这样使用
	var i int=0
	var j int=len(numbers)-1
	for i<j{
		if numbers[i]+numbers[j]==target{
			return []int{i,j}
		}

		if numbers[i]+numbers[j]<target{
			i++
		}else{
			j--
		}
	}
	// xcrj 创建空slice
	return []int{}
}