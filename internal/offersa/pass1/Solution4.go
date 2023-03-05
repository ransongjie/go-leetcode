package pass1

/**
 * 剑指 Offer II 004. 只出现一次的数字
 * 给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
 */

type Solution4 struct{}

    /**
     * 散列表统计每个数字出现的次数
     */
func (solution4 Solution4) singleNumber1(nums []int) int{
	var amap map[int]int=make(map[int]int)
	for num:=range nums{
		// get or default, 不存在t=0
		var t int=amap[num]
		amap[num]=t+1
	}

	// 查找出现1次的值
	for k,v:=range amap{
		if v==1{
			return k
		}
	}

	return -1
}

     /**
     * 数组中只有出现1次或3次的数字，每个数字的二进制表示，每1位取值为0或1
     * 所有数字二进制表示的 第i位之和%3=出现1次数字二进制表示 第i位值
     */
func (solution4 Solution4) singleNumber2(nums []int) int{
	//求出现1次数字的所有32位bit值，求出现1次数字第i位bit值，求所有数字第i位之和%3，求单个元素第i位的值
	// xcrj 指定 int 精确类型
	var r int32=0
	for i:=0;i<32;i+=1{
		var sum int32=0
		// 
		for _,num:=range nums{
			sum+=((int32(num)>>i)&1)
		}

		// 出现1次数字，第i位为1
        // xcrj 将结果值第i位置为1, 1左移i位后 与原结果 按位或
		if sum%3==1{
			r|=(1<<i)
		}
	}

	return int(r)
}