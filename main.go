package main

import (
	"fmt"
	"math"
)

func main() {

	i := 0.0
	fmt.Scan(&i)

	if float64(int(math.Log2(i))) == math.Log2(i) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
