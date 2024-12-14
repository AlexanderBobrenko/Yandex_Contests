package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	get_stolb(a)
}
func get_stolb(n int) {
	itog := ""
	a := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for n > 0 {
		n--
		itog = string(a[n%26]) + itog
		n /= 26
	}
	fmt.Print(itog)
}
