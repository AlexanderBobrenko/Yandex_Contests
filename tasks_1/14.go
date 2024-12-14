package main

import "fmt"

func main() {
	var strok, stolb int
	fmt.Scan(&strok, &stolb)
	fmt.Print("    |")
	for i := 1; i <= stolb; i++ {
		fmt.Printf("%4d", i)
	}
	fmt.Println()
	fmt.Print("   --")
	for j := 0; j < stolb; j++ {
		fmt.Print("----")
	}
	fmt.Println()
	for l := 1; l <= strok; l++ {
		fmt.Printf("%4d|", l)
		for k := 1; k <= stolb; k++ {
			fmt.Printf("%4d", k*l)
		}
		fmt.Println()
	}
	fmt.Print("")
}
