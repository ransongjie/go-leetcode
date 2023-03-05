package pass1

import(
	"strings"
)

/**
 * 剑指 Offer II 005. 不含相同字符的单词的长度的最大乘积
 * - 给定一个字符串数组words，1个元素1个字符串
 * - 计算当两个字符串不包含相同字符时，长度乘积的最大值
 * - 字符串中只包含英语的小写字母
 */

type Solution5 struct{}


    /**
     * 每一个单词和后面的所有单词对比，蛮力法
     * 单词i从0到len
     * 单词j从i到len
     * 
     * @param words
     * @return
     */
func (solution5 Solution5) maxProduct (words []string)int{
	// 不存在满足条件的解 则返回0
	var r int=0
	for i:=0;i<len(words);i+=1{
		var wordi string=words[i]
		for j:=i;j<len(words);j+=1{
			var wordj string=words[j]
			// xcrj go成员方法调用
			if !solution5.hasSameChar1(wordi,wordj) {
				var l int = len(wordi)*len(wordj)
				if l>r{
					r=l
				}
			}
		}
	}

	return r
}

    /**
     * 单词i中是否含有单词j的字符，遍历单词j的每个字符，判断字符是否在单词i中
     * 
     * @param wordi
     * @param wordj
     * @return
     */
func (solution5 Solution5) hasSameChar1 (wordi string, wordj string) bool{
	for _,c:=range wordj{
		if strings.ContainsRune(wordi,c) {
			return true
		}
	}

	return false
}

    /**
     * 利用散列表存储单词中的字符
     * 散列函数 c-'a'=idx
     * 
     * @param wordi
     * @param wordj
     * @return
     */
func (solution5 Solution5) hasSameChar2 (wordi string, wordj string) bool{
	var cs []int=make([]int,26)
	// 存储wordi
	for _,c:=range wordi{
		// xcrj go字符数据类型rune
		var idx rune=c-'a'
		cs[idx]=1
	}

	// 判断wordj的字符是否已经存储到cs散列表中
	for _,c:=range wordj{
		var idx rune=c-'a'
		if 1==cs[idx] {
			return true
		}
	}

	return false
}

    /**
     * 位散列表，两个整数分别存储两个字符串，整数&运算==0 判断是否存在相同字符
     * 散列函数 bm|=1<<(c-'a')
     * 
     * @param wordi
     * @param wordj
     * @return
     */
func (solution5 Solution5) hasSameChar3 (wordi string, wordj string) bool{
	var bm1 int32=0
	var bm2 int32=0
	for _,c:=range wordi{
		bm1|=1<<(c-'a')
	}

	for _,c:=range wordj{
		bm2|=1<<(c-'a')
	}

	if 0==(bm1&bm2){
		return false
	}
	return true
}