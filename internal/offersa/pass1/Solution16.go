package pass1
/**
 * 剑指 Offer II 016. 不含重复字符的最长连续子字符串
 * - 子字符串连续
 * - 子字符串中不含有重复字符
 */
type Solution16 struct{}
// 空结构体作map的value，空结构体不占用空间
type void struct{}
var vo void
	/**
     * 双while夹击+set去重，
     * 先往set中添加元素，
     * 再判断set中是否有重复元素，
     * 有从左往右依次丢弃
     * @param s
     * @return
     */
func (solution16 *Solution16) lengthOfLongestSubstring(s string) int {
	//
	var maxLen int=0
	//string  to []rune
	var cs []rune=[]rune(s)
	var i int=0
	var j int=0
	// xcrj 使用map创建 set
	var set map[rune]void=make(map[rune]void)
	for j<len(s){
		// 代替 while(i<j&&set.contains(s.charAt(j)))
		for i<j{
			_,ok:=set[cs[j]]
			if ok{
				// xcrj 内置函数写法 len() delete()
				delete(set,cs[i])
				i++
			}else{
				break
			}
		}
		set[cs[j]]=vo
		j++
		maxLen=solution16.max(maxLen,j-i)
	}
	return maxLen
}

func (solution16 *Solution16) max(a int,b int)int{
	if a>b{
		return a
	}
	return b
}