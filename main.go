package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(strconv.FormatInt(760, 2))
	fmt.Print(BinGap(760))
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
