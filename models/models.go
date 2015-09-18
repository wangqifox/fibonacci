package models

func FibonacciSequence(number int) int {
	a, b := 0, 1
	for i := 0; i < number; i++ {
		a, b = b, a+b
		if a < 0 {
			return -1
		}
	}
	return a
}
