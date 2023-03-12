package pass1
import (
	"unicode"
)
/**
 * 剑指 Offer II 018. s是否为回文串
 * - 只考虑字母和数字字符
 * - 忽略字母的大小写。
 */
func isPalindrome(s string) bool {
	var i int=0
	var j int=len(s)-1
	var cs []rune=[]rune(s)
	for i<j {
		// !(digit(a)||letter(a))=!digit(a)&&!letter(a)
		for i<j&&!unicode.IsDigit(cs[i])&&!unicode.IsLetter(cs[i]) {i++}
		for i<j&&!unicode.IsDigit(cs[j])&&!unicode.IsLetter(cs[j]) {j--}
		if unicode.ToLower(cs[i])!=unicode.ToLower(cs[j]){
			return false
		}else{
			i++
			j--
		}
	}
	return true
}