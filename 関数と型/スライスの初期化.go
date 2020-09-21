package main

import "fmt"

func main() {
	// ゼロ値はnil
	var ns1 []int
	fmt.Println(ns1)

	// 長さと容量を指定して初期化
	// 各要素はゼロ値で初期化される
	ns1 = make([]int, 3, 10)
	fmt.Println(ns1)

	// スライスリテラルで初期化
	// 要素数は指定しなくてよい
	// 自動で配列は作られる
	var ns2 = []int{10, 20, 30, 40, 50}
	fmt.Println(ns2)

	ns3 := []int{5: 50, 10: 100}
	fmt.Println(ns3)
}
