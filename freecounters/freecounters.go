package main

import "fmt"

func counter() func() {
	ctr := 0
	fmt.Println("カウンタを初期化しました")
	return func() {
		ctr++
		fmt.Println(ctr)
	}
}

func freeCounter(start int) func(int) int {
	ctr := start

	fmt.Printf("カウンタを%dから始めます\n", ctr)
	return func(add int) int {
		fmt.Printf("%dを足して", add)
		ctr += add
		return ctr
	}
}

func main() {
	freeCnt := freeCounter(10)
	fmt.Println(freeCnt(2))
	fmt.Println(freeCnt(5))
	fmt.Println(freeCnt(7))
}
