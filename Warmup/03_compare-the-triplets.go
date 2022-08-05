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
1. Read and understand :
- read trough the problem out loud to helps you slow down
-- clarify any part of it you do not understand
--- ask if you cant ask a question or can you googled by yourself

- what are the possible inputs-outputs? check the arguments that our functions will take?
-- how much var are the expected inputs? is it number or char?
--- write down the expected input-output examples

2. Step-by-step human plan :
- break down the core problem into smaller step, think about how you would solve it as a human
-- write the step-by-step plan (pseudocode) how to solve the problem
--- create a variable to store the expected output with the given input, execute the main process (core solution)


3. Step-by-step machine execution :
- get out your simple, mechanical core solution and put into programming syntax
-- temporalily ignore the hard part, start from the easiest that you can start writing
--- remember not to prematurely optimize, working code first, efficiency later

4. Revise the code :
- can you understand the solution at a glance? does it makes sense?
-- can you derive the result differently? what other approaches are there?
--- can you remove some variable to make the code more simple?
*/

// given func named "compareTriplets" as human problem
// with "a" and "b" as an input param with a type of slice-int32 format
// and a result type of slice-int32
func compareTriplets(a []int32, b []int32) []int32 {
	// below is machinie solution
	// create variable to store value for every output-result comparisons conditions
	var alicePoints, bobPoints int32
	// iterate every index of a
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] { // when input "a" or alice work got higher score than bob's work
			alicePoints += 1 // alice got 1 point
		} else if a[i] < b[i] { // when input "b" or bob work got higher score than alice
			bobPoints += 1 // bob got 1 point
		}
	}
	return []int32{alicePoints, bobPoints} // expected output : 2 slice of array
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)
	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32
	for i := 0; i < 3; i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	bTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var b []int32
	for i := 0; i < 3; i++ {
		bItemTemp, err := strconv.ParseInt(bTemp[i], 10, 64)
		checkError(err)
		bItem := int32(bItemTemp)
		b = append(b, bItem)
	}

	result := compareTriplets(a, b)
	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")
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
