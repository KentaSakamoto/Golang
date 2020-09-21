package main

import "fmt"

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func main() {
	// 100をHex型として代入
	var hex Hex = 100
	// メソッド式
	f := Hex.String
	fmt.Printf("%T\n%s\n", f, f(hex))
}
