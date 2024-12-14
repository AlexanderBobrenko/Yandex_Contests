package main

import (
	"fmt"
	"math"
)

func main() {
	vozd := 365 * 0.5
	dub := 20.0
	topol := 32.0
	fmt.Print(vozd, math.Ceil(vozd/topol), math.Ceil(vozd/dub))
}
