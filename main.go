package main

import ("fmt")

func main() {
	var a float64
	var b float64
	var c float64

	fmt.Print("a = ")
	fmt.Scan(&a)
	fmt.Print("b = ")
	fmt.Scan(&b)
	fmt.Print("c = ")
	fmt.Scan(&c)

	fmt.Println("V =", a * b * c)
	fmt.Println("S =", 2 *( a * b + b * c + a * c))

}
