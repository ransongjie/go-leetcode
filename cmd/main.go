package main
import (
	"fmt"
	"sort"
	"bytes"
	"reflect"
	"com.xcrj.goleetcode/internal/offersa/pass1"
	"strings"
	"unicode"
)
func main() {
	test26()
}
func test27(){
	//slice 实现 队列
}
func test26(){
	//slice 实现 栈
	var limit int=10
	var stack []int=make([]int,0)
	//push 压栈
	for i:=0;i<3;i++{
		stack=append(stack,i)
	}
	//pop 弹出
	var val int=stack[len(stack)-1]
	stack=stack[:len(stack)-1]
	fmt.Println(val)
	//peak 查看栈顶元素
	var pk int=stack[len(stack)-1]
	fmt.Println(pk)
	//栈空
	if len(stack)==0{
		fmt.Println("栈空")
	}
	//栈满
	if len(stack)>limit{
		fmt.Println("栈满")
	}else{
		fmt.Println("栈未满")
	}
}
func test25(){
	//判断字母 数字
	var a rune='a'
	var ab bool=unicode.IsLetter(a)
	fmt.Println(ab)

	var a1 rune='2'
	var ab1 bool=unicode.IsDigit(a1)
	fmt.Println(ab1)
}
func test24(){
	//unicode
	//字母 大写转小写
	var a rune='A'
	var al rune=unicode.ToLower(a)
	fmt.Printf("%c",al)
	//字母 大写转小写 反之
	var a1 rune='a'
	var al1 rune=unicode.ToUpper(a1)
	fmt.Printf("%c",al1)
}
func test23(){
	// strings
	//字符串 大写转小写
	var a string="ABCD"
	var al string=strings.ToLower(a)
	fmt.Println(al)
	//字符串 大写转小写 反之
	var a1 string="abcd"
	var al1 string=strings.ToUpper(a1)
	fmt.Println(al1)
}

func test22(){
	//获取变量类型
	fmt.Println(reflect.TypeOf(pass1.MinWindow("a","a")).String())
}
func test21(){
	// []rune切片转string
	var str1 string="xcrj"
	// 写法 强制类型转换
	var cs []rune=[]rune(str1)
	// 写法 强制类型转换
	var str2 string=string(cs)
	fmt.Println(str2)
}

func test20(){
	// 是否存在
	var am map[string]int=map[string]int{"xcrj":1}
	v:=am["xcrj"]
	fmt.Println(v)
	_,ok:=am["xcrj"]
	fmt.Println(ok)
}

//空结构体不占用空间
type void struct{}
var vv void

func test19(){	
	//创建set
	// map[string]类型
	var set map[string]void=make(map[string]void)
	// set["xcrj"]=值
	set["xcrj"]=vv
	set["xcrj1"]=vv
	fmt.Println(set)
}

func test18(){
	// array to slice 
	var as[3]int=[3]int{1,2,3}
	var s1[]int=as[:]
	fmt.Println(s1)
	// slice to array
	var bs[3]int=[3]int{}
	copy(bs[:],s1)// copy(目的,原)
	fmt.Println(bs)
}

func test17(){
	// 数组判断相等
	var a1 [3]int=[3]int{1,2,3}
	var a2 [3]int=[3]int{1,2,3}
	fmt.Println(a1==a2)

	// byte切片判断相等
	var bs1 []byte=[]byte{1,2,3}
	var bs2 []byte=[]byte{1,2,3}
	fmt.Println(bytes.Equal(bs1,bs2))

	// 普通切片判断相等
	var s1 []int=[]int{1,2,3}
	var s2 []int=[]int{1,2,3}
	// fmt.Println(s1==s2)
	// 反射性能消耗大
	fmt.Println(reflect.DeepEqual(s1,s2))
	// 自己写判断
	var len int=len(s1)
	var eql bool=true
	for i:=0;i<len;i++{
		if(s1[i]!=s2[i]){
			eql=false
			break
		}
	}
	fmt.Println(eql)

}
func test16(){
	// string to []rune
	var str string="xcrj"
	var cs []rune=[]rune(str)
	var c rune=cs[0]
	// 输出字符
	fmt.Printf("%c", c)
}
func test15(){
	var astr string="xcrj"
	fmt.Println(len(astr))
}
func test14(){
	var test14 Test14=Constructor("xcrj")
	test14.ShowName()
}

type Test14 struct{
	// public属性
	Name string
}

// 工厂模式创建类，静态工厂
// public 方法
func Constructor(name string) Test14{
	return Test14{name}
}

func (this *Test14) ShowName(){
	fmt.Println(this.Name)
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
	// slice len 和 cap
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
