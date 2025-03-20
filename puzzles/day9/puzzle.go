package day9

import (
	"aoc2024/internal/utils"
	"fmt"
	"slices"
	"strconv"
)

var emptySpace = "."

func Solve(testFile string, params map[string]interface{}) {
	var inputFile = fmt.Sprintf("puzzles/day9/%s", testFile)
	inputContent, _ := utils.ReadFile(inputFile)

	decodedInput := decodeInput(inputContent)
	decodedInputPart1 := make([]string, len(decodedInput))
	copy(decodedInputPart1, decodedInput)
	moveBlocks(&decodedInputPart1)
	fmt.Println(calcChecksum(decodedInputPart1))

	decodedInputPart2 := make([]string, len(decodedInput))
	copy(decodedInputPart2, decodedInput)
	moveFiles(&decodedInputPart2)
	fmt.Println(calcChecksumWithEmptySpaces(decodedInputPart2))

}

func decodeInput(input string) []string {
	decodedInput := make([]string, 0)
	for i, v := range input {
		multiplier, _ := strconv.Atoi(string(v))

		if i%2 == 0 {
			for j := 0; j < multiplier; j++ {
				decodedInput = append(decodedInput, strconv.Itoa(i/2))
			}
		} else {
			for j := 0; j < multiplier; j++ {
				decodedInput = append(decodedInput, emptySpace)
			}
		}
	}
	return decodedInput
}

func moveBlocks(decodedInput *[]string) {
	for i, v := range *decodedInput {
		if v == emptySpace {
			for j, x := range slices.Backward(*decodedInput) {
				if i >= j {
					return
				}
				if x != emptySpace {
					(*decodedInput)[i] = x
					(*decodedInput)[j] = emptySpace
					break
				}
			}
		}
	}
}

func moveFiles(decodedInput *[]string) {
	// create list of free spaces, with start/end index and length
	freeSpacesList := make([]IndexTruple, 0, len(*decodedInput))
	startIdx := -1
	endIdx := -1
	for i := 0; i < len(*decodedInput); i++ {
		if (*decodedInput)[i] == emptySpace {
			if i > 0 && (*decodedInput)[i-1] == emptySpace {
				endIdx = i
			} else {
				startIdx = i
				endIdx = i
			}
		} else {
			if i > 0 && (*decodedInput)[i-1] == emptySpace {
				spaceLength := endIdx - startIdx + 1
				newSpace := IndexTruple{
					startIdx: startIdx,
					endIdx:   endIdx,
					length:   spaceLength,
				}
				freeSpacesList = append(freeSpacesList, newSpace)
			}
		}
	}

	// iterate over decoded input (backwards)
	// try to find first free space (from left) that matches whole file (from right)
	// whole file - sequence of the same id numbers
	startIdx = -1
	endIdx = -1
	for i := len(*decodedInput) - 1; i >= 0; i-- {
		if (*decodedInput)[i] != emptySpace {
			// first fileId (from the right)
			if i == len(*decodedInput)-1 {
				endIdx = i
				startIdx = i
				continue
			}
			// next (counting from right), fileId same as previous
			if (*decodedInput)[i+1] == (*decodedInput)[i] {
				startIdx = i

				// different fileId than previous or previous was emptySpace
			} else {
				// startIdx = i+1 (shouldn't have to set it manually)

				// different fileId, but previous was NOT emptySpace
				// move previous file if possible
				if (*decodedInput)[i+1] != emptySpace {
					spaceLength := endIdx - startIdx + 1
					value := (*decodedInput)[startIdx]
					moveFile(decodedInput, &freeSpacesList, startIdx, endIdx, spaceLength, value)
				}
				// start checking new file
				endIdx = i
				startIdx = i

			}
		} else {

			// if current input == emptySpace, it's not first from right and previous wasn't emptySpace
			// move previous fileId if possible
			if i < len(*decodedInput)-1 && (*decodedInput)[i+1] != emptySpace {
				spaceLength := endIdx - startIdx + 1
				value := (*decodedInput)[startIdx]
				moveFile(decodedInput, &freeSpacesList, startIdx, endIdx, spaceLength, value)

			}
		}
	}
}

func moveFile(decodedInput *[]string, freeSpaces *[]IndexTruple,
	startIdx, endIdx, spaceLength int, value string) {
	for i, v := range *freeSpaces {
		if v.length >= spaceLength && v.endIdx < startIdx {
			for j := 0; j < spaceLength; j++ {
				(*decodedInput)[v.startIdx+j] = value
			}
			for j := startIdx; j <= endIdx; j++ {
				(*decodedInput)[j] = emptySpace
			}
			// not sure which is faster, and if second one is safe
			if v.length == spaceLength {
				*freeSpaces = slices.Delete(*freeSpaces, i, i+1)
			} else {
				(*freeSpaces)[i].startIdx += spaceLength
				(*freeSpaces)[i].length -= spaceLength
			}
			//*freeSpaces = append((*freeSpaces)[:i], (*freeSpaces)[i+1:]...)
			return
		}
	}

	return
}

// IndexTruple was tuple, but added length, so its `truple` now
type IndexTruple struct {
	startIdx, endIdx, length int
}

func calcChecksum(blocks []string) int {
	blocksToSum := slices.Delete(blocks, slices.Index(blocks, emptySpace), len(blocks))
	checksum := 0
	for i, block := range blocksToSum {
		blockValue, _ := strconv.Atoi(block)
		checksum += i * blockValue
	}
	return checksum
}

func calcChecksumWithEmptySpaces(blocks []string) int {
	checksum := 0
	for i, block := range blocks {
		if block == emptySpace {
			continue
		}
		blockValue, _ := strconv.Atoi(block)
		checksum += i * blockValue
	}
	return checksum
}
