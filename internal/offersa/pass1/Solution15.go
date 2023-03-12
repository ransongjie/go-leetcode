package pass1
/**
 * 剑指 Offer II 015. 字符串中的所有变位词的左索引
 * - p在s串内的变位词的做索引
 */
type Solution15 struct{}
    /**
     * 方法与剑指 Offer II 014. 字符串中的变位词 一致
     * @param s
     * @param p
     * @return
     */
func(solution15 *Solution15) findAnagrams(s string, p string) []int {
	var s1 string=p
	var s2 string=s
	var ls []int=[]int{}
	//特殊情况 xcrj len(string)
	var len1 int=len(s1)
	var len2 int=len(s2)
	if len1>len2{
		return ls
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
	if solution15.eq(cnt1[:],cnt2[:]){
		// xcrj slice 增加 内置函数
		ls=append(ls,0)
	}
	//窗口滑动
	for i:=len1;i<len2;i++ {
		cnt2[cs2[i]-'a']++
		cnt2[cs2[i-len1]-'a']--
		if solution15.eq(cnt1[:],cnt2[:]){
			ls=append(ls,i-len1+1)
		}
	}

	return ls
}

func (solution15 *Solution15) eq(as []int,bs []int)bool{
	var eq bool=true
	for i:=0;i<26;i++{
		if as[i]!=bs[i]{
			eq=false
			break
		}
	}
	return eq
}