package pass1


/**
 * 剑指 Offer II 003. 前n个数字二进制中1的个数
 * 给定一个非负整数 n ，请计算 0 到 n 之间的每个数字的二进制表示中 1 的个数，并输出一个数组。
 * <p>
 * 考点：
 * 发现规律（动态规划）
 * r&1奇偶特性：r为奇数结果为1，r为偶数结果为0
 * r>>1奇偶特性：r为奇数或偶数结果相等，结果都是除以2取整
 */

type Solution3 struct{}

    /**
     * 奇数二进制1个数=(奇数-1)二进制1个数+1
     * 偶数二进制1个数=(偶数/2)二进制1个数
     * 动态规划
     * 奇数：dp[2k+1]=dp[2k]+1=dp[k]+1
     * 偶数：dp[2k]=dp[k]
     * 
     * @param n
     * @return
     */
func (solution3 Solution3) countBits1(n int) []int{
	// 动态规划数组，从0~n共n+1个元素
	// var rs [n]int=[n]int{} 数组必须给准确长度
	var rs []int=make([]int,n+1)
	// 初始状态，0的二进制中1的个数为0
	rs[0]=0
	for i:=1;i<=n;i+=1{
		// 使用动态规划数组
		if i%2==0{
			rs[i]=rs[i>>1]
		}else{
			rs[i]=rs[i>>1]+1
		}
	}

	return rs
}

    /**
     * xcrj i&1 替代 奇偶判断后+0或+1
     * i&1奇偶特性：i为奇数结果为1，i为偶数结果为0
     * @param n
     * @return
     */
func (solution3 Solution3) countBits2(n int) []int{
	// 动态规划数组，从0~n共n+1个元素
	// var rs [n]int=[n]int{} 数组必须给准确长度
	var rs []int=make([]int,n+1)
	// 初始状态，0的二进制中1的个数为0
	rs[0]=0
	for i:=1;i<=n;i+=1{
		rs[i]=rs[i>>1]+(i&1)
	}

	return rs
}