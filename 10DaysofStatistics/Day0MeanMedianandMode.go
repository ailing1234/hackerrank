package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	inputReader := bufio.NewReader(os.Stdin)

	n, inputStringX := GetNInput(inputReader)

	sumOfX, arrayXInt := GetArrayX(inputStringX)

	mean := float64(sumOfX) / float64(n)
	fmt.Printf("%.1f\n", mean)

	median := CalcMedian(arrayXInt, n)
	fmt.Printf("%.1f\n", median)

	mode := CalculateMode(arrayXInt)
	fmt.Printf("%d", mode)
}

func CalculateMode(arrayXInt []int) (mode int) {
	sort.Ints(arrayXInt)
	// Create a map to count occurrences of each element
	frequency := make(map[int]int)
	for _, num := range arrayXInt {
		frequency[num]++
	}

	// Find the mode (element with the highest frequency)
	maxFrequency := 0
	for num, freq := range frequency {
		if freq > maxFrequency {
			maxFrequency = freq
			mode = num
		}
	}

	// If there are ties (multiple modes), choose the smallest numerically
	for num, freq := range frequency {
		if freq == maxFrequency && num < mode {
			mode = num
		}
	}

	return mode

}

func CalcMedian(arrayXInt []int, n int) (median float64) {
	sort.Ints(arrayXInt)

	// if n is odd
	if n%2 == 1 {
		median = float64(arrayXInt[n/2])
		return median
	}

	// If n is even
	median = (float64(arrayXInt[n/2-1]) + float64(arrayXInt[n/2])) / 2.0
	return median
}

// get array X from User
func GetArrayX(inputStringX string) (sumOfX int, arrayXInt []int) {
	inputStringX = strings.TrimSpace(inputStringX)
	arrayX := strings.Fields(inputStringX)
	for _, x := range arrayX {
		xInt, err := strconv.Atoi(x)
		if err != nil {
			fmt.Println("Error in COnversion Atoi")
		}
		sumOfX += xInt
		arrayXInt = append(arrayXInt, xInt)
	}

	return sumOfX, arrayXInt
}

// get input N from user
func GetNInput(inputReader *bufio.Reader) (n int, inputStringX string) {
	inputReader = bufio.NewReader(os.Stdin)
	inputN, _ := inputReader.ReadString('\n')
	inputN = strings.TrimSpace(inputN)

	// convert string to integer
	n, err := strconv.Atoi(inputN)
	if err != nil {
		fmt.Println("Error in COnversion Atoi")
	}

	// check the constrains
	if n < 10 || n > 2500 {
		fmt.Println("N must be greater equal than 10 and less equal than 2500.")
	}

	inputStringX, _ = inputReader.ReadString('\n')
	inputStringX = strings.TrimSpace(inputStringX)

	return n, inputStringX

}
