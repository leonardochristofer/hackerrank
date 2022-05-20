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

 * The Fibonacci numbers are a sequence of numbers where each number after the first two is a sum of the prior two.

 * As an illustration, here is a short sequence given starting values of (0, 1): Fibonacci series =  (0, 1, 1, 2, 3, 5, 8, 13).

 * Given an integer n, calculate the first n numbers in the Fibonacci sequence given starting elements of (0, 1).

 * Return an array of n integers, including the given (0, 1) in the sequence.

 * Function Description: Complete the function fibonacci in the editor below.

 * fibonacci has the following parameter(s):

 * int n:  the length of the Fibonacci series to return

 * Return: int[n]: an array of n Fibonacci numbers starting with (0, 1)

 * Complete the 'fibonacci' function below.

 * The function is expected to return an INTEGER_ARRAY.

 * The function accepts INTEGER n as parameter.

 * Sample Input

 * STDIN   Function
 * -----   --------
 * 4     → n = 4

 * Sample Output

 * 0
 * 1
 * 1
 * 2

 */

func fibonacci(n int32) []int32 {
	// Write your code here
	a := int32(0)
	b := int32(1)
	c := int32(0)
	var fibonacciSlice []int32
	fibonacciSlice = append(fibonacciSlice, 0)
	for i := int32(0); i < n-int32(1); i++ {
		c = a
		a = b
		b = c + a
		fibonacciSlice = append(fibonacciSlice, a)
	}
	return fibonacciSlice
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

	result := fibonacci(n)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
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
