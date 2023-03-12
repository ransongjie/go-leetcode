package pass1
/**
 * 剑指 Offer II 019. 最多删除一个字符得到回文
 * - 从字符串中最多删除一个字符(删除0个字符或1个字符)
 * - 能否得到一个回文字符串
 */

     /**
     * 双指针相向移动
     * - 相等继续移动
     * - 不相等 跳过左字符判断相等回文串||跳过右字符判断回文串
     */
func validPalindrome(s string) bool {
	var i int=0
	var j int=len(s)-1
	var cs []rune=[]rune(s)
	for i<j {
		if cs[i]==cs[j] {
			i++
			j--
		}else{
			return validPalindrome2(s,i+1,j)||validPalindrome2(s,i,j-1)
		}
	}
	return true
}

// xcrj golang 不支持方法重载
// func validPalindrome(s string,start int,end int) bool {
func validPalindrome2(s string,start int,end int) bool {
	var cs []rune=[]rune(s)
	for start<end{
		if cs[start]!=cs[end]{
			return false
		}else{
			start++
			end--
		}
	}
	return true
}