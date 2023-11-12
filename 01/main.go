package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	readFile, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	elfArr := make([]int64, 1)
	elfIndex := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			elfIndex++
			elfArr = append(elfArr, 0)
			continue
		}

		var intval, err = strconv.ParseInt(line, 0, 64)
		if err != nil {
			fmt.Printf("Could not parse integer value %v", line)
			panic("Could not parse integer value")
		}

		elfArr[elfIndex] += intval
	}

	readFile.Close()

	sort.Slice(elfArr, func(i int, j int) bool {
		return elfArr[i] < elfArr[j]
	})

	var topsum int64 = 0
	arrlen := len(elfArr)
	for n := arrlen - 3; n < arrlen; n++ {
		fmt.Println(elfArr[n])
		topsum += elfArr[n]
	}

	fmt.Printf("Total of top 3: %v\n", topsum)
}
