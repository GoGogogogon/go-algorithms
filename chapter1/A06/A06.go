package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var n, q int

	// fmt.Scan(&n)
	// fmt.Scan(&q) fmt.scanは遅い

	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n, &q)
	a := make([]int, n+1) //人数のスライス
	a[0] = 0

	for i := 1; i <= n; i++ {
		var x int
		fmt.Fscan(in, &x)
		a[i] = a[i-1] + x
	}

	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(in, &l, &r)
		fmt.Println(a[r] - a[l-1])
	}

	// l := make([]int, q+1)
	// r := make([]int, q+1)

	// for i := 1; i <= q; i++ {
	// 	fmt.Scan(&l[i], &r[i])

	// }  計算するだけならいらない

	// sum := make([]int, n+1)
	// sum[0] = 0
	// //累積和
	// for i := 1; i <= n; i++ {
	// 	sum[i] = sum[i-1] + a[i]
	// }

	// 11 46 47 77 80
	//  0  11 57 104 181 261

	// s5 = 181 + 80

	// q　質問数
	// for i := 1; i <= q; i++ {
	// 	total := sum[r[i]] - sum[l[i]-1]
	// 	//r日目　- (l-1)日目
	// 	fmt.Println(total)
	// }

}

// func main() {
// 	var d int
// 	var n int

// 	fmt.Scan(&d)
// 	fmt.Scan(&n)

// 	l := make([]int, n+1)
// 	r := make([]int, n+1)

// 	///累積和　b
// 	b := make([]int, d+2)
// 	b[0] = 0

// 	for i := 1; i <= n; i++ {
// 		fmt.Scan(&l[i], &r[i])
// 	}

// 	//前日比に加算
// 	for i := 1; i <= n; i++ {
// 		b[l[i]] += 1
// 		b[r[i]+1] -= 1
// 	}

// 	ans := make([]int, d+2)
// 	ans[0] = 0

// 	for i := 1; i <= d; i++ {
// 		ans[i] = ans[i-1] + b[i]
// 	}
// 	for i := 1; i <= d; i++ {
// 		fmt.Println(ans[i])
// 	}
// 	//出力

// }

// func main() {
// 	var d int
// 	var n int

// 	fmt.Scan(&d)
// 	fmt.Scan(&n)
// 	l := make([]int, n+1)
// 	r := make([]int, n+1)

// 	events := make([]int, d+2)

// 	for i := 1; i <= n; i++ {
// 		fmt.Scan(&l[i], &r[i])
// 	}

// 	// l日目の値　+ 1  r+1 の値 -1

// 	for i := 1; i <= n; i++ {
// 		events[l[i]] += 1
// 		events[r[i]+1] -= 1
// 	}

// 	ans := make([]int, d+2)

// 	for i := 1; i <= d; i++ {
// 		ans[i] += ans[i-1] + events[i]
// 		fmt.Println(ans[i])
// 	}

// }
