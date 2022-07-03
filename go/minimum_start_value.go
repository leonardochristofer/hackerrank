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

 * Start with a given array of integers and an arbitrary initial value x.

 * Calculate the running sum of x plus each array element, from left to right.

 * The running sum must never get below 1.

 * Determine the minimum value of x.

 * Example

 * arr = [-2, 3, 1, -5].

 * If x = 4, the following results are obtained:

 * Running
 * sum       arr[i]
 * -----     -----
 * 4          -2
 * 2           3
 * 5           1
 * 6          -5
 * 1

 * The final value is 1, and the running sum has never dropped below 1.  The minimum starting value for x is 4.

 * Function Description

 * Complete the function minX in the editor below.

 * minX has the following parameter(s):

 * int arr[n]:  an array of integers

 * Returns

 * int: the minimum integer value for x

 * Input Format for Custom Testing

 * Input from stdin will be processed as follows and passed to the function.

 * The first line contains an integer n, the size of the array arr.

 * Each of the next n lines contains an integer arr[i].

 * Sample Case 0

 * Sample Input

 * STDIN     Function
 * -----     -----
 * 10     →  arr[i] size n = 10
 * -5     →  arr = [-5, 4, -2, 3, 1, -1, -6, -1, 0, 5]
 * 4
 * -2
 * 3
 * 1
 * -1
 * -6
 * -1
 * 0
 * 5

 * Sample Output

 * 8

 * Explanation

 * Running
 * sum       arr[i]
 * -----     -----
 * 8          -5
 * 3           4
 * 7          -2
 * 5           3
 * 8           1
 * 9          -1
 * 8          -6
 * 2          -1
 * 1           0
 * 1           5
 * 6

 * The minimum starting value for x is 8.

 * Sample Case 1

 * Sample Input

 * STDIN     Function
 * -----     -----
 * 5      →  arr[i] size n = 5
 * -5     →  arr = [-5, 4, -2, 3, 1]
 * 4
 * -2
 * 3
 * 1

 * Sample Output

 * 6

 * Explanation

 * Running
 * sum      arr[i]
 * -----    -----
 * 6         -5
 * 1          4
 * 5         -2
 * 3          3
 * 6          1
 * 7

 * The minimum starting value for x is 6.

 * Sample Case 2

 * Sample Input

 * STDIN     Function
 * -----     -----
 * 10     →  arr[i] size n = 10
 * -5     →  arr = [-5, 4, -2, 3, 1, -1, -6, -1, 0, -5]
 * 4
 * -2
 * 3
 * 1
 * -1
 * -6
 * -1
 * 0
 * -5

 * Sample Output

 * 13

 * Explanation

 * Running
 * sum        arr[i]
 * -----      -----
 * 13          -5
 * 8           4
 * 12          -2
 * 10           3
 * 13           1
 * 14          -1
 * 13          -6
 * 7          -1
 * 6           0
 * 6          -5
 * 1

 * The minimum starting value for x is 13.

 * Complete the 'minX' function below.

 * The function is expected to return an INTEGER.

 * The function accepts INTEGER_ARRAY arr as parameter.

 */

func minX(arr []int32) int32 {
	// Write your code here
	sum := int32(0)
	minVal := int32(1)
	for _, val := range arr {
		sum += val
		if minVal > sum {
			minVal = sum
		}
	}
	return 1 + (-1 * minVal)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	arrCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var arr []int32

	for i := 0; i < int(arrCount); i++ {
		arrItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := minX(arr)

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
