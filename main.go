package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	val := strings.Split(scanner.Text(), " ")
	num1 := new(big.Int)
	num2 := new(big.Int)
	num3 := new(big.Int)

	_, success1 := num1.SetString(val[0], 10)
	_, success2 := num2.SetString(val[1], 10)
	_, success3 := num3.SetString(val[2], 10)

	if !success1 || !success2 || !success3 {
		return
	}

	max := num1
	if num2.Cmp(max) == 1 {
		max = num2
	}
	if num3.Cmp(max) == 1 {
		max = num3
	}

	fmt.Println(max)

}
