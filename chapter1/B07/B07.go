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

	var T int
	var n int

	fmt.Scan(&T)
	fmt.Scan(&n)
	l := make([]int, n+1)
	r := make([]int, n+1)
	sum := make([]int, T+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&l[i], &r[i])
	}

	//時間の累積和
	for i := 1; i <= n; i++ {
		sum[l[i]] += 1
		sum[r[i]] -= 1
	}

	part := make([]int, T+1)

	part[0] = sum[0]
	fmt.Println(part[0])

	// for i := 1; i <= T-1; i++ {
	// 	fmt.Println(sum[i])
	// }
	for t := 1; t <= T-1; t++ {
		part[t] = part[t-1] + sum[t]
		fmt.Println(part[t])
	}

}
