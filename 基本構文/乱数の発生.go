package main

import (
	"math/rand"
	"time"
)

func main() {
	// 現在時刻を数値で取得する
	t := time.Now().UnixNano()
	// 乱数のたねを設定
	rand.Seed(t)
	// xは0-10の間の値になる
	s := rand.Intn(10)
	println(s)
}
