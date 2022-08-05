package Hackerrank_Problem_Solving

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

package main

import (
"bufio"
"fmt"
"io"
"os"
"strconv"
"strings"
)

// func named "simpleArraySum"
// with "ar" as an input parameter with type of one dimmensional slice in "[]int32" format
// and a result type of int32
func simpleArraySum(ar []int32) int32 {
	// create variable "sum" to store sum of one dimmensional slice value
	var sum int32
	// iterate every index of the "ar" slice []int32
	for i := 0; i < len(ar); i++ {
		// create the sum of value everytime the loop runs sum +=
		sum += ar[i]
	}
	// return the sum value as result
	return sum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	arCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	arTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var ar []int32

	for i := 0; i < int(arCount); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := simpleArraySum(ar)

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
