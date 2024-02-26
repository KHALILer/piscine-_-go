package piscine

func ConcatParams(args []string) string {
	var result string
	size := len(args)

	for i, v := range args {
		result += v

		if i != size-1 {
			result += "\n"
		}
	}
	return result
}
