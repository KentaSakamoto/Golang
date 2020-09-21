package main

func main() {
	ns := []int{10, 20, 30, 40, 50}
	// 要素にアクセス
	println(ns[3])
	// 長さ
	println(len(ns))
	// 要素の追加
	// 容量が足りない場合は背後の配列が再確保される
	ns = append(ns, 60, 70)
	// 容量
	println(cap(ns))

}
