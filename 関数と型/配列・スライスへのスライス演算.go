package main

import "fmt"

func main() {
	ns := []int{10, 20, 30, 40, 50}
	n, m := 2, 4

	// n番目以降のスライスを取得する
	fmt.Println(ns[n:]) // [30 40 50]

	// 先頭からm-1番目までのスライスを取得する
	fmt.Println(ns[:m]) // [10 20 30 40]

	// capを指定する
	ms := ns[:m:m]
	fmt.Println(cap(ms)) // 4
}
