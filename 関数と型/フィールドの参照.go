package main

func main() {
	p := struct {
		name string
		age  int
	}{name: "Gopher", age: 10}
	// フィールドにアクセスする例
	p.age++
	println(p.name, p.age)
}
