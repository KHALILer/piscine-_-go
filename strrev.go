package main

func StrRev(s string) string {
	x := []rune(s)
	l := len(x)
	for i := 0; i < l/2; i++ {
		x[i], x[l-i-1] = x[l-i-1], x[i]
	}
	return string(x)
}
