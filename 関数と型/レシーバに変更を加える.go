package main

type MyInt int
// まだできていない
func (n MyInt) Inc() { n++ }
func main() {
	var n MyInt
	println(n)
	n.Inc()
	println(n)
}
