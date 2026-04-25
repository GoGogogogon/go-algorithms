package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)

	dp := make([]int, n+1)

	suuti := make([]int, n+1)
	//どこから来たかを記録する配列

	a := make([]int, n+1)
	b := make([]int, n+1)

	//データ読み込み
	for i := 2; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}
	for i := 3; i <= n; i++ {
		fmt.Fscan(in, &b[i])
	}

	//初期値
	dp[1] = 0
	dp[2] = a[2]
	suuti[2] = 1
	//suuti[2]は部屋1からくるから1.

	for i := 3; i <= n; i++ {
		costA := dp[i-1] + a[i]
		costB := dp[i-2] + b[i]

		if costA <= costB {
			dp[i] = costA
			suuti[i] = i - 1
		} else {
			dp[i] = costB
			suuti[i] = i - 2
		}
	}

	// ゴールからの逆算
	current := n
	route := make([]int, 0)
	// ルートを記録するスライス

	// 現在地が1（スタート）に辿り着くまで回し続ける
	for current > 1 {

		route = append(route, current)
		if current == 1 {
			break
		}

		current = suuti[current]
	}

	for i := len(route) - 1; i >= 0; i-- {
		fmt.Print(route[i], " ")
	}
	fmt.Println() // 最後に改行
}
