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
	var n int

	fmt.Scanln(&n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	res := strings.Split(input, " ")
	tasks := make([]int, n)

	for i, val := range res {
		num, err := strconv.Atoi(val)
		if err != nil {
			return
		}
		tasks[i] = num
	}
	first := tasks
	slices.Sort(first)

	val := make([]int, 3)
	num := make([]int, 3)
	for i, j := 0, n-1; i < n; i++ {
		if val[2]+tasks[i] <= 300 {
			val[2] += tasks[i]
			num[2] += 1 //5
		}
		if val[1]+tasks[j] <= 300 {
			val[1] += tasks[j]
			num[1] += 1 //3
		}
		if val[0]+first[i] <= 300 {
			val[0] += first[i]
			num[0] += 1 //1
		}

		j--
	}
	temp_num := slices.Index(num, slices.Max(num))
	temp_val := slices.Index(val, slices.Min(val))

	fmt.Println("Tasks:", first)
}
