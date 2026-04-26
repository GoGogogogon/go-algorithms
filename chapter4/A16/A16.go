package main

import (
	"bufio"
	"fmt"
	"os"
)

// func main() {
// 	var n int
// 	in := bufio.NewReader(os.Stdin)

// 	fmt.Fscan(in, &n)

// 	dp := make([]int, n+1)

// 	a := make([]int, n+10)

// 	b := make([]int, n+10)
// 	for i := 2; i <= n; i++ {
// 		fmt.Fscan(in, &a[i])
// 	}

// 	for i := 3; i <= n; i++ {
// 		fmt.Fscan(in, &b)
// 	}
// 	dp[1] = 0
// 	dp[2] = a[2]

// 	for i := 3; i <= n; i++ {
// 		dp[i] = min(dp[i-1]+a[i], dp[i-2]+b[i])
// 	}

// 	fmt.Println(dp)

// }

func main() {
	var n int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &n)

}
