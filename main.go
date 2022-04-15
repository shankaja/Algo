package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Print(Primes(10001))
}

// project-euler 7
func Primes(r int) int {
	c := 1
	p := 0
	for i := 2; c <= r; i++ {
		if i == 2 {
			p = i
			c++
			continue
		}
		d := false
	innerLoop:
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				d = true
				break innerLoop
			}
		}
		if !d {
			p = i
			c++
		}
	}
	return p
}

// project-euler 6
func SumOfSquareDiff(n int) int {
	x := 0
	y := 0
	for i := 1; i <= n; i++ {
		x += int(math.Pow(float64(i), 2))
		y += i
	}
	return (int(math.Pow(float64(y), 2)) - x)
}

// projectEuler q-5
func EvenlyDevidable(m int) int {
	for i := m; true; i += m {
		b := true
	innerLoop:
		for j := 1; j <= m; j++ {
			if i%j != 0 {
				b = false
				break innerLoop
			}
		}
		if b {
			return i
		}
	}
	return 0
}

// ProjectEuler question - 4
func LargestPalindromProduct(d float64) (int, int) {
	// e.g. d= 2-> 10^2-1 = 99 -> 99*99
	x := math.Pow(10, d)
	m := math.Pow((x - 1), 2)
	intX := int(x)
	fmt.Printf("X value : %f \n", x)
	for i := m; i >= 0; i-- {
		if IsPalindrome(int(i)) {
			fmt.Printf("Palindrome found %f \n", i)
			for j := int(x); j > 0; j-- {
				r := int(i) / j
				if int(i)%j == 0 && r < intX {
					return j, r
				}
			}
		}
	}

	return 0, 0
}

func IsPalindrome(n int) bool {
	s := fmt.Sprintf("%d", n)
	charArr := []rune(s)
	l := len(s)
	for i := l / 2; i >= 0; i-- {
		if charArr[i] != charArr[l-i-1] {
			return false
		}
	}
	return true
}

func LargestPrimeFactor(n int) int {
	for i := n / 13; i > 1; i-- {
		if n%i == 0 {
			if IsPrime(i) {
				return i
			}
		}
	}
	return 1
}

func IsPrime(n int) bool {
	d := false
	if n%2 == 0 || n%3 == 0 || n%5 == 0 || n%7 == 0 || n%11 == 0 || n%13 == 0 {
		return false
	}
	for j := 2; j <= n/2; j++ {
		if n%j == 0 {
			d = true
			break
		}
	}
	fmt.Println(n)
	return !d
}

func PrimeSeries(r int) []int {
	ps := make([]int, 0)
	c := 0
	for i := 2; c < r; i++ {

		if c >= r {
			return ps
		}
		if i == 2 {
			ps = append(ps, i)
			c++
			continue
		}
		d := false
	innerLoop:
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				d = true
				break innerLoop
			}
		}
		if !d {
			ps = append(ps, i)
		}
		c++
	}
	return ps
}

func EvenFib(r int) int {
	c := 0
	fv := 0
	for n := 0; fv < r; n++ {
		fv = optimizedFib(n)
		if fv > r {
			continue
		}
		if fv%2 == 0 {
			c += fv
		}
	}
	return c
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

func SumOfMultipliers(n int, m int, r int) int {
	c := 0
	for i := 0; i < r; i++ {
		if i%n == 0 || i%m == 0 {
			c += i
		}
	}
	return c
}

func BinGap(N int) int {
	x := strings.TrimRight(string(strconv.FormatInt(int64(N), 2)), "0")
	charArr := []rune(x)
	maxC := 0
	c := 0
	for _, v := range charArr {
		if string(v) == "1" {
			if maxC < c {
				maxC = c
			}
			c = 0
		} else {
			c++
		}
	}
	return maxC
}

// func Solution(A []int) int {
// 	// write your code in Go 1.4
// 	t := 0
// 	for _, v := range A {
// 		t += v
// 	}
// 	_, c := sol(A, 0, t)
// 	return c
// }

// func sol(A []int, c int, total int) ([]int, int) {
// 	sort.Ints(A)
// 	t := 0
// 	for _, v := range A {
// 		t += v
// 	}
// 	if t <= total/2 {
// 		return A, c
// 	}
// 	A[len(A)-1] = A[len(A)-1] / 2
// 	c++
// 	return sol(A, c, total)

// }

// func Solution(P []int, S []int) int {
// 	length := len(P)
// 	sort.Sort(sort.Reverse(sort.IntSlice(S)))

// 	total := 0
// 	for i := 0; i < length; i++ {
// 		total += P[i]
// 	}
// 	fmt.Printf("Total pax : %d \n", total)
// 	fmt.Printf("Seat availability : %d \n", S)

// 	c := 0
// 	for _, a := range S {
// 		if total-a <= 0 {
// 			fmt.Printf("Total : %d Last Car avl : %d \n Returning!\n", total, a)
// 			return c + 1
// 		}
// 		total -= a
// 		c++
// 	}
// 	return c
// }

// }

// func Solution(message string, K int) string {
// 	length := len(message)
// 	if length <= K {
// 		return strings.TrimRight(message, " ")
// 	} else {
// 		charArr := []rune(message)

// 		for i := K; i >= 0; i-- {
// 			if unicode.IsSpace(charArr[i]) {
// 				return getSlicedStr(charArr[0:i])
// 			}

// 		}
// 	}
// 	return ""
// }

// func getSlicedStr(charArr []rune) string {
// 	op := ""
// 	for _, k := range charArr {
// 		op += string(k)
// 	}
// 	return strings.TrimRight(op, " ")
// }
