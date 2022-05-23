#!/bin/python3

# Task

# The provided code stub reads two integers from STDIN, a and b. Add code to print three lines where:

# The first line contains the sum of the two numbers.

# The second line contains the difference of the two numbers (first - second).

# The third line contains the product of the two numbers.

# Example

# a = 3
# b = 5

# Print the following:

# 8
# -2
# 15

# Input Format

# The first line contains the first integer, a.
# The second line contains the second integer, b.

# Sample Input

# 3
# 2

# Sample Output

# 5
# 1
# 6

# Explanation

# 3 + 2 => 5
# 3 - 2 => 1
# 3 * 2 => 6

import math
import os
import random
import re
import sys

if __name__ == '__main__':
    a = int(input())
    b = int(input())
    print(a + b)
    print(a - b)
    print(a * b)