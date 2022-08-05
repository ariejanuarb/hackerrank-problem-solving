package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) { // input : slice with fixed-5 elements with positive values
	// task: find the minimum and maximum values made from four of five integer values

	// 1. sort arr slice by value from min to max
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j]})

	// 2. store the sum value to variable
	// 3. iterate to sum the value (index is ommited, cause its already sorted)
	// 3.1 requirement: for minSum = iterate range between (0 index) to (max-1 index),
	var minSum int
	for _, value := range arr[:len(arr)-1] { //iterate start from range[lowbound=0:highbound=4] = [0,1,2,3,4] = [0,1,2,3]
		minSum += int(value) // store the added values per loop
	}

	// requirement: for maxSum = iterate range between (1 index) to (max index)
	var maxSum int
	for _, value := range arr[1:] { //iterate start from range[lowbound=1:highbound=max=5]  = [0,1,2,3,4] = [1,2,3,4]
		maxSum += int(value)
	}

	// output : single line of two space-separated long int = {minSum, maxSum}
	// 4. print minSum and maxSum in the same line
	fmt.Println(minSum, maxSum)
}

sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j]}) //sort by value from min to max

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
