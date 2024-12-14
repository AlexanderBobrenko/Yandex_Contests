package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]float64, 0, n)
	for i := 0; i < n; i++ {
		var a float64
		fmt.Scan(&a)
		s = append(s, a)
	}
	for ind, _ := range s {
		if ind == 0 || ind == len(s)-1 {
			fmt.Print(s[ind], " ")
			continue
		} else {
			fmt.Print((s[ind-1]+s[ind]+s[ind+1])/3, " ")
		}
	}
}
