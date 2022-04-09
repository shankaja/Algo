package main

import (
	"fmt"
	"time"
)

func main() {
	steps := 50
	start := time.Now()
	ways := calculateUsingFib(steps)
	elapsed := time.Since(start)
	fmt.Printf("Time elapsed for fibbonnaci solution : %f answer %d \n", elapsed.Seconds(), ways)

	start = time.Now()
	ways = calculateUsingOptimizedFib(steps)
	elapsed = time.Since(start)
	fmt.Printf("Time elapsed for optimized fibbonnaci solution : %f answer %d", elapsed.Seconds(), ways)
}

func calculateUsingFib(n int) int {
	return fib(n + 1)
}

// this basic fibbonnaci sequence algorithm needs to be optimized
// time complexity of this would be O(2^n) which is a fairly high complexity
func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func calculateUsingOptimizedFib(n int) int {
	return optimizedFib(n + 1)
}

func optimizedFib(n int) int {
	if n >= 0 && n < 2 {
		return n
	} else {
		b := make([]int, 0)
		b = append(b, 0, 1)
		for i := 2; i < n+1; i++ {
			b = append(b, b[i-1]+b[i-2])
		}
		return b[n]
	}
}
