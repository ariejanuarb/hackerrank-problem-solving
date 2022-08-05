package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func plusMinus(arr []int32) { // the input was slice with index number of n-len

	// create variable to store the count result value of 3 type number
	var countPositive, countNegative, countZero int
	// iterate every value of slice with condition to count and add the point
	for _, value := range arr { // omit index "_," only count value
		if value > 0 {
			countPositive++
		} else if value < 0 {
			countNegative++
		} else {
			countZero++
		}
	}
	// output should return 3 lines : 1.count of positive value/n; 2.count of negative val./n; 3. count of zero value/n --with 6 digit decimals
	// with 6 decimal place round float
	fmt.Println(float64(countPositive) / float64(len(arr))) // it should be Println, otherwise the rounding will change
	fmt.Println(float64(countNegative) / float64(len(arr)))
	fmt.Println(float64(countZero) / float64(len(arr)))

}

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
