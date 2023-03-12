package pass1
/**
 * 剑指 Offer II 014. s1是s2字符串的变位词
 * - s1的变位词（每个排列）在s2中
 * - s1长度大于s2
 */
type Solution14 struct{}
    /**
     * 排列，每种字符的数量相等，次序可以不同
     * 固定滑动窗口+散列表统计
     * - int[] cnt1s=new int[26];
     * - 窗口长度为s1的长度
     * 
     */
func(solution14 *Solution14) checkInclusion(s1 string, s2 string) bool {
	//特殊情况 xcrj len(string)
	var len1 int=len(s1)
	var len2 int=len(s2)
	if len1>len2{
		return false
	}
	//初始化 xcrj 数组创建 {}
	var cnt1 [26]int=[26]int{}
	var cnt2 [26]int=[26]int{}
	// xcrj 字符串转字符切片
	var cs1 []rune=[]rune(s1)
	var cs2 []rune=[]rune(s2)
	for i:=0;i<len1;i++{
		cnt1[cs1[i]-'a']++
		cnt2[cs2[i]-'a']++
	}
	// xcrj array 转 slice
	if solution14.eq(cnt1[:],cnt2[:]){
		return true
	}
	//窗口滑动
	for i:=len1;i<len2;i++ {
		cnt2[cs2[i]-'a']++
		cnt2[cs2[i-len1]-'a']--
		if solution14.eq(cnt1[:],cnt2[:]){
			return true
		}
	}

	return false
}

func(solution14 *Solution14) eq(as []int,bs []int)bool{
	var eq bool=true
	for i:=0;i<26;i++{
		if as[i]!=bs[i]{
			eq=false
			break
		}
	}
	return eq
}