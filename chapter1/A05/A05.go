package main

import "fmt"

func main() {

	var n int
	var k int
	var count int
	fmt.Scan(&n)
	fmt.Scan(&k)

	//3 6

	// 1000 6000
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			m := k - i - j

			if m >= 1 && m <= n {
				count++
				//fmt.Println("i, j, mの値", i, j, m)
			}

		}
	}
	fmt.Println(count)
}
