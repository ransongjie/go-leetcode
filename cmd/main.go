package main
import (
	"fmt"
	"sort"
)
func main() {

}

func test13(){
	var as []int=[]int{}
	// xcrj go 只有有序列表 没有 无序列表 set
	// 重复值 后者 不会 覆盖前者 有序列表
	as=append(as,1)
	as=append(as,1)
	// [1,1]
	fmt.Println(as)
}

func test12(){
	var as []int=[]int{}
	as=append(as,1)
	as=append(as,3)
	as=append(as,2)
	// 不需要 as=sort.Ints(as)
	sort.Ints(as)
	fmt.Println(as)
}
func test11(){
	var c rune='b'-'a'
	var s []int=make([]int,2)
	s[c]=1
	fmt.Println(s)
}

func test10(){
	var str string="abc"
	for _,c:=range str{
		fmt.Printf("%c", c)
	}
}

func test9(){
	
	var a int=1
	fmt.Println(a<<1)

	// 最好使用精确类型
	var b int32=1
	fmt.Println(b<<1)
}

func test8(){
	// 创建空slice 必须写0
	var as []int=make([]int,0)
	as=append(as,4)
	as=append(as,5)
	as=append(as,6)

	// 必须 for _,a:=range as
	for a:= range as{
		// 0 1 2, 默认输出索引值
		fmt.Println(a)
	}
}

func test7(){
	var amap map[string]int=make(map[string]int)
	amap["xcrj"]=1
	v,ok:=amap["xcrj1"]
	fmt.Println(v," ",ok)
	// get or default
	var r int=amap["xcrj1"]
	fmt.Println(r)
	amap["xcrj1"]=r
	fmt.Println(amap)
}

func test6(){
	var amap map[string]int=make(map[string]int,2)
	// 0, 实际长度
	fmt.Println(len(amap))
	amap["xcrj"]=1
	amap["xcrj2"]=2
	amap["xcrj3"]=3
	// 3, 实际长度
	fmt.Println(len(amap))
	// no cap func for map
	// fmt.Println(cap(amap))
}

func test5(){
	var a string="a"
	var b string ="b"
	fmt.Println(a+b)
}
func test4(){
	var as []rune=make([]rune,0)
	as=append(as,'1')
	var bs []rune=make([]rune,0)
	bs=append(bs,'1')
	bs=append(bs,'2')

	var ras []rune=make([]rune,len(bs)-len(as))
	// a slice 插入 b slice 末尾
	ras=append(ras,as...)
	// [0 49]
	fmt.Println(ras)
}

func test3(){
	var as []rune=make([]rune,0)
	as=append(as,'1')
	var bs []rune=make([]rune,0)
	bs=append(bs,'1')
	bs=append(bs,'2')

	var ras []rune=make([]rune,len(bs))
	// dest src
	copy(ras,as)
	// [49 0]
	fmt.Println(ras)
	for _,a:=range ras{
		// 格式化为字符后输出
		fmt.Printf("%c",a)
	}
}

func test2(){
	var a []rune=nil
	// []
	fmt.Println(a)
	if 1==1 {
		a=make([]rune,2)
	}
	// [0 0]
	fmt.Println(a)
}

func test1(){
	// 1=len 2=cap
	// a[0]默认值为0，赋值len长度的值为0
	var a []int=make([]int,1,2)
	// 1, 
	fmt.Println(len(a))
	a=append(a,1)
	a=append(a,2)
	// 3, 实际长度
	fmt.Println(len(a))
	// 4, todo(slice cap 如何变化)
	fmt.Println(cap(a))
}
