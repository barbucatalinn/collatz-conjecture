package calculator

// isEven returns true if the given number os even
func isEven(n uint64) bool {
	return n%2 == 0
}

// isOdd returns true if the given number is odd
func isOdd(n uint64) bool {
	return n%2 != 0
}

// calculate returns the number of steps required for a number (positive integer)
// to get the to the "4->2->1 sequence"
// if the number is odd we multiply it by 3 and add 1 (3x+1)
// if the number is even we divide it by 2 (x/2)
func calculate(n uint64, steps int) int {
	steps++

	switch {
	case n == 4:
		steps += 3
	case isEven(n):
		n = n / 2
		return calculate(n, steps)
	case isOdd(n):
		n = (n * 3) + 1
		return calculate(n, steps)
	}

	return steps
}

// CalculateCollatzConjectureSteps is a wrapper for the 'calculate' function
// writes the result into the channel
func CalculateCollatzConjectureSteps(n uint64, ch chan map[uint64]int) {
	ch <- map[uint64]int{n: calculate(n, 0)}
}
