package main

import (
	"bufio"
	"fmt"
	"os"
)

// func main() {
// 	in := bufio.NewReader(os.Stdin)
// 	var n int
// 	fmt.Fscan(in, &n)

// 	dp := make([]int, n+1)

// 	suuti := make([]int, n+1)
// 	//どこから来たかを記録する配列

// 	a := make([]int, n+1)
// 	b := make([]int, n+1)

// 	//データ読み込み
// 	for i := 2; i <= n; i++ {
// 		fmt.Fscan(in, &a[i])
// 	}
// 	for i := 3; i <= n; i++ {
// 		fmt.Fscan(in, &b[i])
// 	}

// 	//初期値
// 	dp[1] = 0
// 	dp[2] = a[2]
// 	suuti[2] = 1
// 	//suuti[2]は部屋1からくるから1.

// 	for i := 3; i <= n; i++ {
// 		costA := dp[i-1] + a[i]
// 		costB := dp[i-2] + b[i]

// 		if costA <= costB {
// 			dp[i] = costA
// 			suuti[i] = i - 1
// 		} else {
// 			dp[i] = costB
// 			suuti[i] = i - 2
// 		}
// 	}

// 	// ゴールからの逆算
// 	current := n
// 	route := make([]int, 0)
// 	// ルートを記録するスライス

// 	// 現在地が1（スタート）に辿り着くまで回し続ける
// 	for current >= 1 {

// 		route = append(route, current)
// 		if current == 1 {
// 			break
// 		}

// 		current = suuti[current]
// 	}

//		fmt.Println(len(route))
//		for i := len(route) - 1; i >= 0; i-- {
//			fmt.Print(route[i], " ")
//		}
//		fmt.Println() // 最後に改行
//	}
func main() {
	var n int
	in := bufio.NewReader(os.Stdin)

	fmt.Fscan(in, &n)

	dp := make([]int, n+1)
	//その地点においての最短コスト

	a := make([]int, n+1)
	b := make([]int, n+1)

	suuti := make([]int, n+1)
	//部屋がどこから来たかの番号 (インデックス)

	for i := 2; i <= n; i++ {
		fmt.Fscan(in, &a[i])
	}

	for i := 3; i <= n; i++ {
		fmt.Fscan(in, &b[i])
	}

	suuti[1] = 0
	suuti[2] = 1

	dp[1] = 0
	dp[2] = a[2]

	for i := 3; i <= n; i++ {
		if dp[i-1]+a[i] <= dp[i-2]+b[i] {
			dp[i] = dp[i-1] + a[i]
			suuti[i] = i - 1

		} else {
			dp[i] = dp[i-2] + b[i]
			suuti[i] = i - 2
		}

		fmt.Println(suuti[i])
	}

	current := n //最終地点

	route := make([]int, 0)
	for current >= 1 {

		route = append(route, current)
		//どんどん今の値をrouteスライスに入れる

		if current == 1 {
			break
		}

		current = suuti[current]
	}

	fmt.Println(len(route))

	for i := len(route) - 1; i >= 0; i-- {
		fmt.Print(route[i], " ")
	}

}
