package main

import "fmt"

func main() {

	// 時刻0 開店
	// 時刻T 閉店

	//出勤li 退勤　ri

	//出勤　+1 退勤  -1

	// t = 0 .1.... t-1
	// t + 0.5
	// t = 0.5 1.5 ... t - 0.5

	//退勤時間 r1 = 3 3.5→いない　= -1

	var t int
	var n int

	fmt.Scan(&t)
	fmt.Scan(&n)
	l := make([]int, n)
	r := make([]int, n)

	for i := 1; i <= n; i++ {
		fmt.Scan(&l[i], r[i])
	}
}
