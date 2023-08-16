package stringutils

// Reverse reverses given string
// by Russ Cox - https://groups.google.com/g/golang-nuts/c/oPuBaYJ17t4/m/PCmhdAyrNVkJ
func Reverse(s string) string {
	r := make([]rune, len(s))

	n := 0
	for _, c := range s {
		r[n] = c
		n++
	}

	r = r[0:n]
	for i := 0; i < n/2; i++ {
		r[i], r[n-1-i] = r[n-1-i], r[i]
	}

	return string(r)
}

// ReverseVigo reverses given string too! a little buggy!
func ReverseVigo(s string) string {
	ss := make([]rune, len(s))

	for i, c := range s {
		ss[len(s)-1-i] = c
	}
	return string(ss)
}
