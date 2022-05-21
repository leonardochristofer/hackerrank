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

 * Example

 * A = [1,2,3]

 * Return [3,2,1]

 * Function Description: Complete the function reverseArray in the editor below.

 * reverseArray has the following parameter(s):

 * int A[n]: the array to reverse

 * Returns: int[n]: the reversed array

 * Input Format

 * The first line contains an integer, N, the number of integers in A.

 * The second line contains N space-separated integers that make up A.

 * Sample Input

 * 4
 * 1 4 3 2

 * Sample Output

 * 2 3 4 1

 * Complete the 'reverseArray' function below.

 * The function is expected to return an INTEGER_ARRAY.

 * The function accepts INTEGER_ARRAY a as parameter.

 */

func reverseArray(a []int32) []int32 {
	// Write your code here
	var reversedArray []int32
	for i := int32(len(a) - 1); i >= 0; i-- {
		reversedArray = append(reversedArray, a[i])
	}
	return reversedArray
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(arrCount); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	res := reverseArray(arr)

	for i, resItem := range res {
		fmt.Fprintf(writer, "%d", resItem)

		if i != len(res)-1 {
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
