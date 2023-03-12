package pass1

func countSubstrings(s string) int {
	var cnt int=0
	var cs []rune=[]rune(s)
	for i:=0;i<len(s)*2-1;i++{
		var l int=i/2
		var r int=l+i%2
		for l>=0&&r<len(s)&&cs[l]==cs[r]{
			cnt++
			l--
			r++
		}
	}
	return cnt
}