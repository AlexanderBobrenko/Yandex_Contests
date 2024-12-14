package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	i := 0
	for {
		if x%2 == 0 {
			if x == 1 {
				break
			}
			i++
			x = x / 2
		} else if x == 1 {
			break
		} else {
			i++
			x = 3*x + 1
		}
	}
	fmt.Print(i)

}
