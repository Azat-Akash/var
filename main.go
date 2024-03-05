package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	salary := make([]int, 3)
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	res := strings.Split(reader.Text(), " ")

	for i := 0; i < 3; i++ {
		val, err := strconv.Atoi(res[i])

		if err != nil {
			return
		}

		salary[i] = val
	}

	fmt.Println(slices.Max(salary) - slices.Min(salary))
}
