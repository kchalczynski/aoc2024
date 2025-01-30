package main

import (
	"bufio"
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
	rules, pages := readFile(inputFile)
	pageOrder := parseRulesToMap(rules)

	correctPages := validatePagesFromMap(pages, pageOrder)

	pretty.Println(pageOrder)
	fmt.Println("---------------------------")
	pretty.Println(correctPages)

	fmt.Println(sumMiddleElements(pages, correctPages))
}

/*func validatePages(pages [][]int, pageOrder []int) []int {
	var correctPages []int

	for i, page := range pages {
		validatePage(page, pageOrder)
		correctPages = append(correctPages, i)
	}

	return correctPages
}*/

/*func validatePage(page []int, pageOrder []int) bool {
	for j := 1; j < len(page); j++ {

		if !(slices.Index(pageOrder, page[j]) > slices.Index(pageOrder, page[j-1])) {
			return false
		}
	}
	return true
}*/

func validatePagesFromMap(pages [][]int, pageOrder map[int][]int) []int {

	correctPages := make([]int, 0)

	for i, page := range pages {
		if validatePage(page, pageOrder) {
			correctPages = append(correctPages, i)
		}
	}

	return correctPages
}

func validatePage(page []int, pageOrder map[int][]int) bool {
	for i := 0; i < len(page); i++ {
		beforePages := pageOrder[page[i]]
		for j := i - 1; j >= 0; j-- {
			if slices.Contains(beforePages, page[j]) {
				return false
			}
		}
	}
	return true
}

func sumMiddleElements(pages [][]int, correctPages []int) int {
	var sum = 0
	for _, page := range correctPages {
		sum += pages[page][len(pages[page])/2]
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

func parseRules(rules []OrderPair) []int {
	ruleList := make([]int, 0)
	for _, rule := range rules {
		fmt.Println(ruleList)
		if !slices.Contains(ruleList, rule.page) {
			if !slices.Contains(ruleList, rule.beforePage) {
				ruleList = append(ruleList, rule.page)
				ruleList = append(ruleList, rule.beforePage)
			} else {
				ruleList = slices.Insert(
					ruleList, slices.Index(ruleList, rule.beforePage), rule.page)
			}
		} else {
			if !slices.Contains(ruleList, rule.beforePage) {
				ruleList = append(ruleList, rule.beforePage)
			} else {
				pageBeforeIndex := slices.Index(ruleList, rule.beforePage)
				pageIndex := slices.Index(ruleList, rule.page)

				// for X | Y <-- X must be before Y
				// delete X from list, insert it at Y index,
				// rest of slice (from Y index) is shifted right
				if pageIndex > pageBeforeIndex {
					ruleList = slices.Delete(ruleList, pageIndex, pageIndex+1)
					ruleList = slices.Insert(ruleList, pageBeforeIndex, rule.page)
				}
			}
		}
	}
	return ruleList
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
		//pages[i] = make([]int, len(page))
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
