package main

import (
	"fmt"
	"log"
)

type stack struct {
	items []int
}

func stack_base() {
	var n int
	var ans int
	var pop int
	var err error
	stack := make([]int, 0)

	for {
		fmt.Println("pushしますか？ yes = 1")
		fmt.Scan(&ans)
		if ans == 1 {
			fmt.Scan(&n)
			stack = stack_push(n, stack)
		}

		fmt.Println("popしますか？ yes = 1 ")

		fmt.Scan(&ans)

		if ans == 1 {
			pop, stack, err = stack_pop(stack)
			log.Print(pop, err)
		}

		fmt.Println("終了しますか？y = -1")
		fmt.Scan(&ans)
		if ans == -1 {
			break
		}

	}

}

func (* s stack) stack_push(n int) {

	stack = append(stack, n)
	return stack

}

func stack_pop(stack []int) (int, []int, error) {
	//後からのものを最初に取り出す

	if len(stack) == 0 {
		return 0, nil, fmt.Errorf("スタックが空です")
	}
	pop := stack[len(stack)-1]
	stack_pop := stack[:len(stack)-1]
	return pop, stack_pop, nil
}

//interfaceの理解用

// type Speker interface {
// 	Speak() string
// }

// //Speaker() というメソッドを持っているというルール（interface)

// type Dog struct{}

// func (d Dog) Speak() string {
// 	return "ワン!"
// }

// func MakeSound(s Speker) {
// 	fmt.Println(s.Speak())
// }

// func main() {
// 	dog := Dog{}
// 	MakeSound(dog)
// }

//anyの理解

// func PrintAnything(value any) {
// 	fmt.Println("中身は", value)
// }

// func main() {
// 	word := "りんご"
// 	number := 3.14159265358979
// 	PrintAnything(word)
// 	PrintAnything(number)
// }

//anyは型チェックを放棄　実務では乱用を避ける

//ジェネリクス　anyの便利さと型チェックの両立

// func PrintTwice[T any](value T) {
// 	fmt.Println(value, value)
// }

// func main() {
// 	word := "りんご"
// 	number := 3.14
// 	PrintTwice(word)
// 	PrintTwice(number)
// }
