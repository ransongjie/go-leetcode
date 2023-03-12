package pass1

import(
	"math"
)
/**
 * 剑指 Offer II 017. s含有t所有字符的连续最短字符串
 * - 变位词+字符数量可以更多
 */

var smap map[rune]int=make(map[rune]int)
var tmap map[rune]int=make(map[rune]int)

func MinWindow(s string, t string) string {
	//特殊情况
	if len(s)<len(t){
		return ""
	}
	//统计t字符数量
	var tc[]rune=[]rune(t)
	for _,c:=range tc{
		tmap[c]=tmap[c]+1
	}
	//双while夹击
	var i int=0
	var j int=-1
	var minLen=math.MaxInt
	var left=0
	var ss []rune=[]rune(s)
	for j<len(s){
		//先添加字符，满足s子串字符数量>=t字符数量
		j++
		if j<len(s){
			_,ok:=tmap[ss[j]]
			if ok{
				smap[ss[j]]=smap[ss[j]]+1
			}
		}
		//再删除字符，满足s子串字符数量<=t字符数量
		for i<=j&&ge(){
			if j-i+1<minLen{
				minLen=j-i+1
				left=i
			}

			_,ok:=smap[ss[i]]
			if ok{
				smap[ss[i]]=smap[ss[i]]-1
			}
			i++
		}
	}

	if minLen==math.MaxInt{
		return ""
	}
	//xcrj []rune 转 string
	var r string=string(ss[left:left+minLen])
	return r
}

func ge()bool{
	// s子串字符数量<t字符数量 returnfalse
	for k,v:=range tmap{
		// xcrj smap[k]不存在返回0
		if smap[k]<v{
			return false
		}
	}

	return true
}