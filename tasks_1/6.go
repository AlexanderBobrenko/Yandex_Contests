package main

import "fmt"

func main() {
	var n, pt, ot, p, dv, st int
	fmt.Scan(&n)
	for n > 0 {
		if n >= 5000 {
			pt += n / 5000
			n %= 5000
			continue
		} else if n >= 1000 {
			ot += n / 1000
			n %= 1000
			continue
		} else if n >= 500 {
			p += n / 500
			n %= 500
			continue
		} else if n >= 200 {
			dv += n / 200
			n %= 200
			continue
		} else if n >= 100 {
			st += n / 100
			n %= 100
			continue
		}
	}
	fmt.Print(pt, ot, p, dv, st)
}
