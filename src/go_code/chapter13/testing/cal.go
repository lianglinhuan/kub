package cal

func Add(n int) int {
	res := 0
	for i := 0 ; i <= n - 1; i++ {
		res = res + i
	}
	return res
}

func Sub(n int, m int) int {
	return n - m
}