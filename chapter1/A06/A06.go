package main

import "fmt"

// func main() {

// 	var n int
// 	var q int

// 	fmt.Scan(&n)
// 	fmt.Scan(&q)
// 	a := make([]int, n+1)

// 	for i := 1; i <= n; i++ {
// 		fmt.Scan(&a[i])

// 	}

// 	l := make([]int, q+1)
// 	r := make([]int, q+1)

// 	for i := 1; i <= q; i++ {
// 		fmt.Scan(&l[i], &r[i])

// 	}

// 	sum := make([]int, n+1)
// 	sum[0] = 0
// 	//累積和
// 	for i := 1; i <= n; i++ {
// 		sum[i] = sum[i-1] + a[i]
// 	}

// 	// 11 46 47 77 80
// 	//  0  11 57 104 181 261

// 	// s5 = 181 + 80

// 	// q　質問数
// 	for i := 1; i <= q; i++ {
// 		total := sum[r[i]] - sum[l[i]-1]
// 		//r日目　- (l-1)日目
// 		fmt.Println(total)
// 	}

// }

func main() {
	var d int
	var n int

	fmt.Scan(&d)
	fmt.Scan(&n)

	l := make([]int, n)
	r := make([]int, n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&l[i], &r[i])
	}

	//出力

}
