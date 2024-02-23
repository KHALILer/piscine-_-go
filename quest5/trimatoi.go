package piscine

func TrimAtoi(s string) int {
	slc := 1
	res := 0
	found := false
	for _, i := range s {
		if i >= '0' && i <= '9' {
			found = true
			res = res*10 + int(i-'0')

		} else if i == '-' && !found {
			slc = -1
		}
	}
	return slc * res
}
