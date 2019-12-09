package common

func Permutations(n int) [][]int {
	if n < 0 {
		panic("n should be non-negative")
	}

	total := factorial(n)

	perms := make([][]int, total)
	for i := 0; i < total; i++ {
		perms[i] = make([]int, n)
	}

	insert := func(a []int, val, pos, l int) {
		for j := l - 1; j > pos; j-- {
			a[j] = a[j-1]
		}
		a[pos] = val
	}

	offset := 1
	for i := 1; i < n; i++ {
		for j, end := 0, offset; j < end; j++ {
			prev := perms[j]
			for k := 0; k < i; k++ {
				copy(perms[offset], prev)
				insert(perms[offset], i, k, n)
				offset++
			}
			insert(prev, i, i, n)
		}
	}
	return perms
}

func factorial(n int) int {
	if n < 0 {
		panic("n should be non-negative")
	}
	r := 1
	for i := 1; i <= n; i++ {
		r *= i
	}
	return r
}
