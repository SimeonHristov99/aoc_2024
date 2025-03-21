package day01

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
	"testing"
)

// Tests that the input files are correctly read and parsed.
func TestParseInput(t *testing.T) {
	t.Run("should read and parse the file 'sample.txt' correctly", func(t *testing.T) {
		// Arrange
		filename := "sample.txt"
		expected := [2][]int{{3, 4, 2, 1, 3, 3}, {4, 3, 5, 3, 9, 3}}

		// Act
		actual, err := ParseInput(filename)

		// Assert
		if err != nil {
			t.Fatalf("Expected successful execution, but got error: %s", err)
		}
		if !reflect.DeepEqual(*actual, expected) {
			t.Fatalf("Expected %v, but got %v.", expected, *actual)
		}
	})

	t.Run("should log out an error when unable to load file", func(t *testing.T) {
		// Arrange
		filename := "invalidSample.txt"
		expectedText := fmt.Sprintf("Could not open file '%s' due to ", filename)

		// Act
		actualResult, actualError := ParseInput(filename)
		actualText := actualError.Error()

		// Assert
		if actualResult != nil {
			t.Fatalf("Expected to return nil for result, but got %v.", actualResult)
		}
		if !strings.Contains(actualText, expectedText) {
			t.Fatalf("Expected error %s, got %s.", expectedText, actualText)
		}
	})
}

func TestSortSlice(t *testing.T) {
	t.Run("should sort correctly", func(t *testing.T) {
		// Arrange
		slice := []int{4, 3, 5, 3, 9, 3}
		expected := []int{3, 3, 3, 4, 5, 9}

		// Act
		actual := SortSlice(slice)

		// Assert
		if !slices.Equal(actual, expected) {
			t.Fatalf("Expected %v, but got %v.", expected, actual)
		}
	})
}

func TestSumPairDiffs(t *testing.T) {
	t.Run("should work correctly for sample", func(t *testing.T) {
		// Arrange
		ls1 := []int{1, 2, 3, 3, 3, 4}
		ls2 := []int{3, 3, 3, 4, 5, 9}
		expected := 11

		// Act
		actual := SumPairDiffs(ls1, ls2)

		// Assert
		if actual != expected {
			t.Fatalf("Expected %d, but got %d.", expected, actual)
		}
	})
}

func TestPart1(t *testing.T) {
	t.Run("should return an error on invalid file", func(t *testing.T) {
		// Arrange
		filename := "invalid.txt"
		expectedText := fmt.Sprintf("Could not open file '%s' due to ", filename)

		// Act
		actualResult, actualError := Part1(filename)
		actualText := actualError.Error()

		// Assert
		if actualResult != -1 {
			t.Fatalf("Expected to return -1 for result, but got %v.", actualResult)
		}
		if !strings.Contains(actualText, expectedText) {
			t.Fatalf("Expected error %s, got %s.", expectedText, actualText)
		}
	})

	t.Run("should work correctly for sample", func(t *testing.T) {
		// Arrange
		filename := "sample.txt"
		expected := 11

		// Act
		actual, err := Part1(filename)

		// Assert
		if err != nil {
			t.Fatalf("Expected successful execution, but got error: %s", err)
		}
		if actual != expected {
			t.Fatalf("Expected %d, but got %d.", expected, actual)
		}
	})

	t.Run("should work correctly for input", func(t *testing.T) {
		// Arrange
		filename := "input.txt"
		expected := 11

		// Act
		actual, err := Part1(filename)

		// Assert
		if err != nil {
			t.Fatalf("Expected successful execution, but got error: %s", err)
		}
		if actual != expected {
			t.Fatalf("Expected %d, but got %d.", expected, actual)
		}
	})
}
