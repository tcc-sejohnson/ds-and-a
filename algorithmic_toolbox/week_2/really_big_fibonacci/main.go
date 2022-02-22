package main

// Goal: Find the remainder of a really, really big Fibonacci number n when divided by m.
// I think the best way to do this actually isn't to memoize and recurse like we usually do -- I think it's to
// naively start counting up until we hit a period.

func main() {

}

func FindPisanoPeriodLength(m int) int {
	var a, b, c int
	a = 0
	b = 1
	c = a + b
	for i := 0; i < m*m; i++ {
		c = (a + b) % m
		a = b
		b = c
		if a == 0 && b == 1 {
			return i + 1
		}
	}
	return 0
}

func FibModulo(n, m int) int {
	pisanoPeriod := FindPisanoPeriodLength(m)
	newN := n % pisanoPeriod

	memo := make(map[int]int)
	return fibModuloRecurse(newN, m, memo)
}

func fibModuloRecurse(n int, m int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if fibNum, exists := memo[n]; exists {
		return fibNum
	}
	fibNum := (fibModuloRecurse(n-1, m, memo) + fibModuloRecurse(n-2, m, memo)) % m
	memo[n] = fibNum
	return fibNum
}
