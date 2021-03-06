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

 * Given an array of integers, find the sum of its elements.

 * For example, if the array ar = [1,2,3], 1 + 2 + 3 = 6, so return 6.

 * Function Description: Complete the simpleArraySum function in the editor below. It must return the sum of the array elements as an integer.

 * simpleArraySum has the following parameter(s):

 * ar: an array of integers

 * Input Format

 * The first line contains an integer, n, denoting the size of the array.

 * The second line contains n space-separated integers representing the array's elements.

 * Output Format

 * Print the sum of the array's elements as a single integer.

 * Sample Input

 * 6
 * 1 2 3 4 10 11

 * Sample Output

 * 31

 * Explanation

 * We print the sum of the array's elements: 1 + 2 + 3 + 4 + 10 + 11 = 31.

 * Complete the 'simpleArraySum' function below.

 * The function is expected to return an INTEGER.

 * The function accepts INTEGER_ARRAY ar as parameter.

 */

func simpleArraySum(ar []int32) int32 {
	// Write your code here
	var countArraySum int32
	for _, value := range ar {
		countArraySum += value
	}
	return countArraySum
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
