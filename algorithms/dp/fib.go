package dp

// memoization
var memo map[int]int = make(map[int]int)

func Fib1(n int) int {
	if result, ok := memo[n]; ok {
		return result
	}
	if n <= 1 {
		memo[n] = n
		return n
	}
	return Fib1(n-1) + Fib1(n-2)
}

// bottom-up dp
func Fib2(n int) int {
	memo := make(map[int]int)
	memo[0] = 0
	memo[1] = 1
	for i := 2; i <= n; i++ {
		memo[n] = memo[n-1] + memo[n-2]
	}
	return memo[n]
}
