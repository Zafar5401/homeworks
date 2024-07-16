package main
import "fmt"

func main() {
	var a int
	var a1 int8
	var a2 int16
	var a3 int32
	var a4 int64
	fmt.Println(a, a1, a2, a3, a4)
	var b uint
	var b1 uint8
	var b2 uint16
	var b3 uint32
	var b4 uint64
	fmt.Println(b, b1, b2, b3, b4)
	var c float32
	var c1 float64
	var c2 complex64
	var c3 complex128
	fmt.Println(c, c1,  c2, c3)
	var d rune
	var d1 byte
	var d2 string
	var d3 bool
	fmt.Println(d, d1, d2, d3)
	a = 4
	a1 = 8
	a2 = 12
	a3 = 16
	a4 = 20
	b = 5
	b1 = 10
	b2 = 15
	b3 = 20
	b4 = 25
	c = 2.5
	c1 = 3.5
	c2 = 4.5
	c3 = 5.5
	d = 30
	d1 = 40
	d2 = "Zafar"
	d3 = true
	fmt.Println(a, a1, a2, a3, a4)
	fmt.Println(b, b1, b2, b3, b4)
	fmt.Println(c, c1, c2, c3)
	fmt.Println(d, d1, d2, d3)
}