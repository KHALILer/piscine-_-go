package piscine

func Capitalize(s string) string {
	result := []rune(s)
	found := true
	for i := range s {
		if (result[i] >= 'a' && result[i] <= 'z') || (result[i] >= 'A' && result[i] <= 'Z') || (result[i] >= '0' && result[i] <= '9') {
			if found {
				if result[i] >= 'a' && result[i] <= 'z' {
					result[i] -= 32
				}
				found = false
			} else {
				if result[i] >= 'A' && result[i] <= 'Z' {
					result[i] += 32
				}
			}
		} else {
			found = true
		}
	}
	return string(result)
}
