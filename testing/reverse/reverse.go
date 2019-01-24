package reverse

//Reverse returns the reverse of the string passed as `s`.
//For example, if `s` is "abcd" this will return "dcba".
func Reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}