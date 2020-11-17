package util

func PageNum(num int) (n int) {
	n = num - 1
	if n < 0 {
		n = 0
	}
	if n > 30 {
		n = 30
	}
	return
}

func Offset(num, total int) (n int) {
	if total < 0 {
		total = 0
	}
	if num < 0 {
		return 0
	}
	if num > total {
		return total
	}
	return
}
