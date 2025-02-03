package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

import "github.com/kr/pretty"

func main() {

	inputFile := "./input.txt"
	outputFile, err := os.Create("./output.txt")
	if err != nil {
		fmt.Println("Failed to create file: ", err)
		return
	}
	defer outputFile.Close()

	rules, pages := readFile(inputFile)
	pageOrder := parseRulesToMap(rules)

	correctPages, incorrectPages := validatePagesFromMap(pages, pageOrder)

	pretty.Println(pageOrder)
	fmt.Println("---------------------------")
	pretty.Println(correctPages)

	fmt.Println(sumMiddleElements(pages, correctPages))
	fmt.Println("---------------------------")

	for _, page := range incorrectPages {
		fmt.Println(page, ":", pages[page])
	}

	/*outputFile.WriteString(pretty.Sprint(pageOrder))
	outputFile.WriteString(fmt.Sprintln("---------------------------"))
	outputFile.WriteString(pretty.Sprint(correctPages))

	outputFile.WriteString(fmt.Sprintln(sumMiddleElements(pages, correctPages)))
	outputFile.WriteString(fmt.Sprintln("---------------------------"))

	for _, page := range incorrectPages {
		outputFile.WriteString(fmt.Sprintln(page, ":", pages[page]))
	}*/

	fmt.Println(sumMiddleElements2(fixIncorrectUpdates(incorrectPages, pages, pageOrder)))

}

/*func validatePage(page []int, pageOrder []int) bool {
	for j := 1; j < len(page); j++ {

		if !(slices.Index(pageOrder, page[j]) > slices.Index(pageOrder, page[j-1])) {
			return false
		}
	}
	return true
}*/

func validatePagesFromMap(pages [][]int, pageOrder map[int][]int) ([]int, []int) {

	correctPages := make([]int, 0)
	incorrectPages := make([]int, 0)

	for i, page := range pages {
		rulesForUpdate := filterRulesByPagesInUpdate(page, pageOrder)
		if validatePage(page, rulesForUpdate) {
			correctPages = append(correctPages, i)
		} else {
			incorrectPages = append(incorrectPages, i)
		}
	}

	return correctPages, incorrectPages
}

func validatePage(page []int, rulesForUpdate map[int][]int) bool {
	for i := 0; i < len(page); i++ {
		beforePages := rulesForUpdate[page[i]]
		for j := i - 1; j >= 0; j-- {
			if slices.Contains(beforePages, page[j]) {
				return false
			}
		}
	}
	return true
}

func filterRulesByPagesInUpdate(page []int, pageOrder map[int][]int) map[int][]int {
	filteredRules := make(map[int][]int)
	for i := 0; i < len(page); i++ {

		filteredRules[page[i]] = getIntArrayIntersection(page, pageOrder[page[i]])
		fmt.Println(page)
		pretty.Println(filteredRules)
	}

	return filteredRules
}

func fixIncorrectUpdates(incorrectPages []int, updates [][]int, pageOrder map[int][]int) [][]int {
	sortedPages := make([][]int, 0)
	for _, page := range incorrectPages {
		rulesForPage := convertRulesToOrder(filterRulesByPagesInUpdate(updates[page], pageOrder))
		sortedPage := sortPagesInUpdate(updates[page], rulesForPage)
		sortedPages = append(sortedPages, sortedPage)
	}
	return sortedPages
}

func convertRulesToOrder(rulesForUpdate map[int][]int) map[int]int {
	rulesPerPageQuantity := make(map[int]int)
	for pageNum, pageRules := range rulesForUpdate {
		rulesPerPageQuantity[pageNum] = len(pageRules)
	}
	return rulesPerPageQuantity
}

func sortPagesInUpdate(update []int, rulesPerPage map[int]int) []int {
	updateSorted := update
	slices.SortFunc(updateSorted, func(i, j int) int {
		return cmp.Compare(rulesPerPage[i], rulesPerPage[j])
	})

	return updateSorted
}

func getIntArrayIntersection(page []int, beforePages []int) []int {
	var result []int

	set := make(map[int]struct{})
	for _, pageNum := range page {
		set[pageNum] = struct{}{}
	}

	for _, pageNum := range beforePages {
		if _, ok := set[pageNum]; ok {
			result = append(result, pageNum)
		}
	}
	return result
}

func sumMiddleElements(pages [][]int, correctPages []int) int {
	var sum = 0
	for _, page := range correctPages {
		sum += pages[page][len(pages[page])/2]
	}
	return sum
}

func sumMiddleElements2(pages [][]int) int {
	var sum = 0
	for i := range pages {
		sum += pages[i][len(pages[i])/2]
	}
	return sum
}

func parseRulesToMap(rules []OrderPair) map[int][]int {
	rulesDict := make(map[int][]int)

	for _, pair := range rules {
		beforePages, ok := rulesDict[pair.page]
		if !ok {
			beforePages = make([]int, 0)
			beforePages = append(beforePages, pair.beforePage)
			rulesDict[pair.page] = beforePages
		} else {
			beforePages = append(beforePages, pair.beforePage)
			rulesDict[pair.page] = beforePages
		}
	}
	return rulesDict
}

func readFile(fileName string) (rules []OrderPair, pages [][]int) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file: ", err)
		return nil, nil
	}

	const (
		RuleSeparator string = "|"
		PageSeparator string = ","
	)

	defer file.Close()
	var fullInput string = ""
	var rulesInput string = ""
	var pageInput string = ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fullInput += scanner.Text() + "\n"
	}
	rulesInput, pageInput = splitInput(splitByEmptyNewline(fullInput))

	for _, line := range strings.Split(rulesInput, "\n") {
		rule := strings.Split(line, RuleSeparator)
		pageNumber, _ := strconv.Atoi(rule[0])
		pageBefore, _ := strconv.Atoi(rule[1])

		rules = append(rules, OrderPair{
			page:       pageNumber,
			beforePage: pageBefore,
		})
	}

	pagesRows := strings.Split(pageInput, "\n")
	pages = make([][]int, len(pagesRows))

	for i, line := range strings.Split(pageInput, "\n") {
		page := strings.Split(line, PageSeparator)
		pages[i] = stringsToInts(page)
	}

	return rules, pages
}

func splitByEmptyNewline(str string) []string {
	strNormalized := regexp.
		MustCompile("\r\n").
		ReplaceAllString(str, "\n")

	return regexp.
		MustCompile(`\n\s*\n`).
		Split(strNormalized, -1)

}

func splitInput(input []string) (string, string) {
	return input[0], input[1]
}

func stringsToInts(input []string) []int {
	result := make([]int, len(input))
	for i, s := range input {
		result[i], _ = strconv.Atoi(s)
	}
	return result
}

type OrderPair struct{ page, beforePage int }
