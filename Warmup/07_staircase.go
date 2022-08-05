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
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func staircase(n int32) { // input : n int32 as last/max size of staircase
	// task : right aligned staircase, composed of # and spaces, has a height and width of n
	// sub task
	for i := 1; int32(i) <= n; i++ {
		// use strings.Repeat("<string value>", <repeated times>)
		fmt.Print(strings.Repeat(" ", int(n)-i)) // right aligned using fmt.Print
		fmt.Println(strings.Repeat("#", i))
		// ketika i = 1 :: print " " * 3, lalu println "#" * 1
	}
}

func main() {
	// read the input from container box sized 1024 * 1024
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	// convert string (input value from box) to int, trim the space (if any)
	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	// check error on value
	checkError(err)
	// n declared an asiigned as int32, parse nTemp from int to int32
	n := int32(nTemp)
	// call the function
	staircase(n)
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
