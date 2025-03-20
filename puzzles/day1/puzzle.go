package day1

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Solve(testFile string, params map[string]interface{}) {

	var col1 []int
	var col2 []int

	col1, col2 = readFileIntoLists(testFile, col1, col2)
	fmt.Println(col1)
	fmt.Println(col2)

	slices.Sort(col1)
	slices.Sort(col2)

	// part 1 solution

	var sum int = 0

	for i := 0; i < len(col1) || i < len(col2); i++ {
		fmt.Println("[", col1[i], "]:[", col2[i], "]")

		if col1[i] >= col2[i] {
			sum += col1[i] - col2[i]
		} else {
			sum += col2[i] - col1[i]
		}
	}
	fmt.Println(sum)

	// part 2 solution

	var similarity int = 0
	occurrenceMap := make(map[int]int)
	for _, element := range col2 {
		if occurrenceMap[element] == 0 {
			occurrenceMap[element] = 1
		} else {
			occurrenceMap[element] += 1
		}
	}

	for _, element := range col1 {
		if occurrenceMap[element] != 0 {
			similarity += element * occurrenceMap[element]
		}
	}

	fmt.Println(similarity)
}

func readFileIntoLists(fileName string, l1 []int, l2 []int) ([]int, []int) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		numbers := strings.Fields(scanner.Text())
		fmt.Println(numbers)

		if len(numbers) > 0 {
			val1, _ := strconv.Atoi(numbers[0])
			val2, _ := strconv.Atoi(numbers[1])
			l1 = append(l1, val1)
			l2 = append(l2, val2)
		}

	}
	return l1, l2
}
