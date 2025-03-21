package day01

import (
	"fmt"
	"math"
	"os"
	"unicode"
)

func ParseInput(filename string) (*[2][]int, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Could not open file '%s' due to %s", filename, err)
	}
	listSwitcher := 0
	var ls1 []int
	var ls2 []int
	for _, r := range string(data) {
		if unicode.IsNumber(r) && listSwitcher == 0 {
			ls1 = append(ls1, int(r-'0'))
			listSwitcher = 1
		} else if unicode.IsNumber(r) && listSwitcher == 1 {
			ls2 = append(ls2, int(r-'0'))
			listSwitcher = 0
		}
	}
	return &[2][]int{ls1, ls2}, nil
}

func SortSlice(slice []int) []int {
	for i := 0; i < len(slice)-1; i++ {
		currMin := i
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[currMin] {
				currMin = j
			}
		}
		slice[currMin], slice[i] = slice[i], slice[currMin]
	}
	return slice
}

func SumPairDiffs(ls1 []int, ls2 []int) int {
	sumDiff := 0
	for i := 0; i < len(ls1); i++ {
		fmt.Printf("%v\t%v\t%d\t%d\n", float64(ls1[i]), float64(ls2[i]), int(math.Abs(float64(ls1[i]) - float64(ls2[i]))), sumDiff)
		sumDiff += int(math.Abs(float64(ls1[i]) - float64(ls2[i])))
	}
	return sumDiff
}

func Part1(filename string) (int, error) {
	slices, err := ParseInput(filename)
	if err != nil {
		return -1, err
	}
	ls1 := SortSlice(slices[0])
	ls2 := SortSlice(slices[1])
	return SumPairDiffs(ls1, ls2), nil
}
