package pass1

import (
	"fmt"
)

/**
 * 剑指 Offer II 002. 二进制加法
 * 给定两个 01 字符串 a 和 b ，请计算它们的和，并以二进制字符串的形式输出。
 * 输入为 非空 字符串且只包含数字 1 和 0。
 * 字符串如果不是 "0" ，就都不含前导零。
 * <p>
 * 进制加法：
 * 二进制逢二进一，商为进位，余数为当前值
 */

type Solution2 struct{}

func (solution2 Solution2) addBinary(a string,b string)string{
	// 处理"0"字符串
	if "0"==a{
		return b
	}
	if "0"==a{
		return a
	}
	// 字符串转储到字符数组，并右对齐
	// 创建空 char slice
	var ras []rune=nil
	var rbs []rune=nil
	// 字符串转字符数组
	var as[]rune=[]rune(a)
	var bs[]rune=[]rune(b)
	if len(as)==len(bs){
		ras=as
		rbs=bs
	}
	if len(as)>len(bs) {
		ras=as
		rbs=make([]rune,len(as)-len(bs))
		rbs=append(rbs,bs...)
	}
	if len(as)<len(bs) {
		rbs=bs
		ras=make([]rune,len(bs)-len(as))
		ras=append(ras,as...)
	}
	// 字符数组转整数数组
	var ias[]int=make([]int,len(ras))
	var ibs[]int=make([]int,len(rbs))
	for i,a:=range ras{
		switch(a){
		case '0':
			ias[i]=0
		case '1':
			ias[i]=1
		}
	}
	for i,b:=range rbs{
		switch(b){
		case '0':
			ibs[i]=0
		case '1':
			ibs[i]=1
		}
	}
	// 二进制加法，(a+b)/2=商 是进位，(a+b)%2=余数是当前值
    // 可能最高位有进位 所以+1
	var ris []int=make([]int,len(ias)+1)
	// xcrj 从低位计算到高位，倒着求解
	// 多变量for循环，i++是语句
	for i,j:=len(ias)-1,len(ris)-1;i>=0;i,j=i-1,j-1{
		ris[j-1]=(ias[i]+ibs[i]+ris[j])/2
		ris[j]=(ias[i]+ibs[i]+ris[j])%2
	}
	// 判断ris的最高位是否为1，不能输出011...
	if ris[0]==1{
		var s string=""
		for _,r:=range ris{
			b:=fmt.Sprintf("%d",r)
			s=s+b
		}
		return s
	}else{
		var s string=""
		for _,r:=range ris[1:]{
			b:=fmt.Sprintf("%d",r)
			s=s+b
		}
		return s
	}
}