package main

import "fmt"

const (
	u         = iota      // u == 0     (untyped integer constant)  iota = 0
	v float64 = iota * 42 // v == 42.0  (float64 constant)			 iota = 1
	w         = iota      // w == 84    (untyped integer constant)	 iota = 2
	a         = 1 << iota // a == 8							     iota = 3
	b                     // b == 16		1 << a					 iota = 4
	c                     // c == 32		1 << b					 iota = 5
	d = iota              // d == 6								 iota = 6
	e = 23                // e = 23								 iota = 7
	f                     // f = 23								 iota = 8
	g = iota              // g = 9									 iota = 9
)

func main() {

	fmt.Println(u, v, w, a, b, c, d, e, f, g)

	const (
		b  = iota + 1
		kb = 1 << (iota * 10) // left shift 1 by 10 times i.e  00000000 00000001 -> 00000100 00000000 i.e 1 becomes 1024
		mb
		gb
	)

	fmt.Printf("%d\t\t%b\n", b, b)
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)
}
