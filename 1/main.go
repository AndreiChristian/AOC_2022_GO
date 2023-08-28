package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	fileName := "input.txt"

	fileContent, err := ReadFile(fileName)
	if err != nil {
		fmt.Println("error ", err)
	}

	var sums []int

	lines := strings.Split(fileContent, "\n\n")
	for _, line := range lines {

		sum := 0
		for _, l := range strings.Split(line, "\n") {
			v, _ := strconv.Atoi(l)
			sum += v
		}

		sums = append(sums, sum)

	}

	sort.Ints(sums)

	top := sums[len(sums)-4 : len(sums)-1]

	finalSum := 0
	for _, v := range top {
		finalSum += v
	}

	fmt.Println(finalSum)

}
