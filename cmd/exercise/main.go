package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Quadratic Coefficients Solver
func main() {
	// fmt.Println(quarter(0))
	// fmt.Println(isUppercase("A DSFFG"))
	// fmt.Println(evenNum(1298734))
	// fmt.Println(closeCompare(1.990000, 5.000000, 3.000000))
	// fmt.Println(contamination("AsvfDfd", "i"))
	// fmt.Println(evenOdd(3))
	// fmt.Println(arrayCSV([][]int{{0, 1, 2, 3, 4}, {10, 11, 12, 13, 14}, {20, 21, 22, 23, 24}, {30, 31, 32, 33, 34}}))
	fmt.Println(derive(5, 9))
}

/*
Given a month as an integer from 1 to 12, return to which quarter of the year it belongs as an integer number.
For example: month 2 (February), is part of the first quarter; month 6 (June), is part of the second quarter;
and month 11 (November), is part of the fourth quarter.

Constraint:
1 <= month <= 12
*/
func quarter(month int) int {
	var quarterN int
	if 1 <= month && month <= 12 {
		switch month {
		case 1, 2, 3:
			quarterN = 1
		case 4, 5, 6:
			quarterN = 2
		case 7, 8, 9:
			quarterN = 3
		case 10, 11, 12:
			quarterN = 4
		}
	} else {
		return quarterN
	}
	return quarterN
}

/*
Create a method to see whether the string is ALL CAPS.

"c" -> False
"C" -> True
"hello I AM DONALD" -> False
"HELLO I AM DONALD" -> True
"ACSKLDFJSgSKLDFJSKLDFJ" -> False
"ACSKLDFJSGSKLDFJSKLDFJ" -> True
*/

func isUppercase(str string) bool {
	if len(str) == 0 {
		return true
	}

	count := 0
	for _, v := range str {
		if unicode.IsSpace(v) {
			count++
		} else if unicode.IsUpper(v) {
			count++
		}
	}

	if count == len(str) {
		return true
	} else {
		return false
	}
}

/*
Return the Nth Even Number

1 --> 0 (the first even number is 0)
3 --> 4 (the 3rd even number is 4 (0, 2, 4))
100 --> 198
1298734 --> 2597466
*/

func evenNum(num int) int {
	if num == 0 || num == 1 {
		return 0
	} else {
		return (num * 2) - 2
	}
}

/*
Create a function close_compare that accepts 3 parameters: a, b, and an optional margin. The function should return whether a is lower than, close to, or higher than b.

Please note the following:

When a is close to b, return 0.
For this challenge, a is considered "close to" b if margin is greater than or equal to the absolute distance between a and b.
Otherwise...

When a is less than b, return -1.
When a is greater than b, return 1.

If margin is not given, treat it as if it were zero.
Assume: margin >= 0

Tip: Some languages have a way to make parameters optional.

Example 1
If a = 3, b = 5, and margin = 3, then close_compare(a, b, margin) should return 0.
This is because a and b are no more than 3 numbers apart.

Example 2
If a = 3, b = 5, and margin = 0, then close_compare(a, b, margin) should return -1.
This is because the distance between a and b is greater than 0, and a is less than b.

best practice
func CloseCompare(a, b, margin float64) int {
  if math.Abs(a-b) <= margin { return 0 }
  if a > b { return 1 }
  return -1
}
*/

func closeCompare(a, b, margin float64) int {
	if a > b {
		if margin != 0.0 && margin >= a-b {
			return 0
		} else if a < b {
			return -1
		} else if a > b {
			return 1
		}
	} else if a < b {
		if margin != 0.0 && margin >= b-a {
			return 0
		} else if a < b {
			return -1
		} else if a > b {
			return 1
		}
	}
	return 0
}

/*
An AI has infected a text with a character!!

This text is now fully mutated to this character.

Starting with the original text, and given a character, return the text once it has been mutated so that all of the characters in the original text have been replaced with the character.

If the text or the character are empty, return an empty string.
There will never be a case when both are empty as nothing is going on!!

Note: The character is a string of length 1 or an empty string.

text before = "abc"
character   = "z"
text after  = "zzz"
*/

func contamination(text, char string) string {
	const emptyStr = ""
	if len(text) == 0 || len(char) == 0 {
		return emptyStr
	}

	var sb strings.Builder

	for i := 0; i < len(text); i++ {
		sb.WriteString(char)
	}

	return sb.String()
}

/*
Create a function that takes an integer as an argument and returns "Even" for even numbers or "Odd" for odd numbers.
*/
func evenOdd(n int) string {
	if n%2 == 0 {
		return "Even"
	}
	return "Odd"
}

/*
Create a function that returns the CSV representation of a two-dimensional numeric array.

input:
   [[ 0, 1, 2, 3, 4 ],
    [ 10,11,12,13,14 ],
    [ 20,21,22,23,24 ],
    [ 30,31,32,33,34 ]]

output:
     '0,1,2,3,4\n'
    +'10,11,12,13,14\n'
    +'20,21,22,23,24\n'
    +'30,31,32,33,34'

Array's length > 2.
*/

func arrayCSV(array [][]int) string {
	var rows []string

	for _, row := range array {
		var rowStr []string
		for _, num := range row {
			rowStr = append(rowStr, strconv.Itoa(num))
		}
		rows = append(rows, strings.Join(rowStr, ","))
	}

	return strings.Join(rows, "\n")
}

/*
This function takes two numbers as parameters, the first number being the coefficient, and the second number being the exponent.
Your function should multiply the two numbers, and then subtract 1 from the exponent.
Then, it has to return an expression (like 28x^7). "^1" should not be truncated when exponent = 2.

derive(7, 8) --> this should output "56x^7"
derive(5, 9) --> this should output "45x^8"

In this case, the function should multiply 7 and 8, and then subtract 1 from 8. It should output "56x^7", the first
number 56 being the product of the two numbers, and the second number being the exponent minus 1.

The output of this function should be a string
The exponent will never be 1, and neither number will ever be 0
*/
func derive(coefficient, exponent int) string {
	if coefficient == 0 || exponent == 0 || exponent == 1 {
		return ""
	}

	var result string
	if exponent == 2 {
		result = fmt.Sprintf(strconv.Itoa(coefficient*exponent) + "x" + strconv.Itoa(exponent))
	} else {
		result = fmt.Sprintf(strconv.Itoa(coefficient*exponent) + "x" + "^" + strconv.Itoa(exponent-1))
	}

	return result
}
