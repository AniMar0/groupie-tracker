package TRC

func Atoi(num string) int {
	digit := 0
	for _, d := range num {
		if d >= '0' && d <= '9' {
			digit = digit*10 + (int(d) - 48)
		}
	}
	return digit
}
