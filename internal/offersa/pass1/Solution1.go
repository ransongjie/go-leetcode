package pass1

import (
	"math"
)

/**
 * 剑指 Offer II 001. 整数除法
 * 给定两个整数 a 和 b ，求它们的除法的商 a/b ，要求不得使用乘号 '*'、除号 '/' 以及求余符号 '%' 。
 * 整数除法的结果应当截去（truncate）其小数部分
 * - xcrj 使用位移运算符 结果没有小数
 * 假设我们的环境只能存储 32 位有符号整数 数值范围是 [−2^31, 2^31−1]。本题中，如果除法结果溢出，则返回 2^31 − 1
 */
type Solution1 struct{}

func (solution1 Solution1) divide(a int, b int) int {
	//a是最小负数，b是-1，结果超过表示范围，直接返回 2^31 − 1
	if a==math.MinInt32 && b==-1{
		return math.MaxInt32
	}

	//后面将使用负数表示，这里先存储计算结果符号
	var  positive bool=false
	if a<0&&b<0{
		positive=true
	}
	if a>0&&b>0{
		positive=true
	}

	//使用负数表示
	//xcrj int64, mb左移时可能超出int的表示范围，所以使用int64
	var ma int64=int64(a)
	if a>0{
		ma=-int64(a)
	}
	var mb int64=int64(b)
	if b>0{
		mb=-int64(b)
	}
	//被除数<除数 2<3 2/3=0
	if ma>mb{
		return 0
	}

	//ma<=mb*2^n ma-mb*2^n=p p仍然能够整除mb
	//求边界，ma<mb, 先让mb>ma，再让mb从左一步步ma，直到ma<mb
	var r int=0;
	var shift int=31
	for{
		if ma>mb {
			break
		}
		for{
			if ma<=(mb<<shift){
				break
			}
			shift--
		}
		ma-=mb<<shift
		r+=1<<shift
	}

	if positive{
		return r
	}else{
		return -r
	}
}