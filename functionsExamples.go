package main

func f1() int {
	return f2()
}

func f2() int {
	return 1
}

// multiples values
func vals() (int, int) {
	return 12, 13
}
func plus(a, b int) int {
	sum := a + b
	return sum
}

// closure

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
