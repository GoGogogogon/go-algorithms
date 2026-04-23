package main

import "fmt"

// func main() {
// 	var n int
// 	fmt.Scan(&n)

// 	for i := 9; i >= 0; i-- {
// 		xi := 1 << i
// 		count := (n / xi) % 2
// 		fmt.Print(count)
// 		// 0000000000
// 	}
// 	fmt.Println("")

// 	// 1か0かの式を用意し、毎回計算することで実装している
// }

func main() {

	var n string
	var sum int
	fmt.Scan(&n)
	//1100
	for i := 0; i < len(n); i++ {
		if n[i] == '1' {
			a := len(n) - 1 - i
			sum += 1 << a
		}
	}
	fmt.Print(sum)
}
