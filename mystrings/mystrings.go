package mystrings

func Reverse(s string) string {
	b := []rune(s)
	for i := range b {
	  j := len(b) - i - 1
	  b[i], b[j] = b[j], b[i]
	}
	return string(b)
}