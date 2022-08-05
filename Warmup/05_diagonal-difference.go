package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// func named "diagonalDifference"
// with "arr" as an input param with a type of two-dimmensional-slice with int32 format
// and a result type of int32
func diagonalDifference(arr [][]int32) int32 {
	var leftRight, rightLeft int32  // var
	for i := 0; i < len(arr); i++ { // len(arr) = len(3) = 0,1,2 consist 3x3 slice
		leftRight += arr[i][i]            // leftRight = leftright + arr[0][0]; leftright + arr[1][1]; leftright + arr[2][2]
		rightLeft += arr[i][len(arr)-i-1] // rightLeft = rightLeft + arr[0][len(arr)-0-1]; rightLeft + arr[1][len(arr)-0-2]; rightLeft + arr[2][len(arr)-0-3]
	}
	return int32(math.Abs(float64(leftRight) - float64(rightLeft))) // absolute diagonal difference
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var arr [][]int32
	for i := 0; i < int(n); i++ {
		arrRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(n) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := diagonalDifference(arr)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
