// Given an array of integers, calculate the ratios of its elements that are positive, negative, and zero.
// Print the decimal value of each fraction on a new line with places after the decimal.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
}

func plusMinus(arr []int32) {
	arrLength := float64(len(arr))
	var positivesProp, negativesProp, zeroesProp float64
	var positivesCount, negativesCount, zeroesCount float64

	for _, i := range arr {
		switch {
		case i > 0:
			positivesCount++
		case i == 0:
			zeroesCount++
		case i < 0:
			negativesCount++
		}
	}

	positivesProp = positivesCount / arrLength
	positivesPropRes := fmt.Sprintf("%.6f", positivesProp)

	negativesProp = negativesCount / arrLength
	negativesPropRes := fmt.Sprintf("%.6f", negativesProp)

	zeroesProp = zeroesCount / arrLength
	zeroesPropRes := fmt.Sprintf("%.6f", zeroesProp)

	fmt.Println(positivesPropRes)
	fmt.Println(negativesPropRes)
	fmt.Println(zeroesPropRes)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
