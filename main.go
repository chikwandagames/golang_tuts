package main

func main() {
	r := factorial(4)
	println(r)

	j := factorial2(4)
	println(j)

	l := factorial3(4)
	println(l)
}

// RECURSION
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factorial2(n int) int {
	val := 1
	for i := n; i > 0; i-- {
		val = val * i
	}
	return val
}

func factorial3(n int) int {
	val := 1
	for ; n > 0; n-- {
		val *= n
	}
	return val
}
