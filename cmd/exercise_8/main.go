package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
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
	// fmt.Println(derive(5, 9))
	// fmt.Println(zooTails("abs", '\u0073'))
	// fmt.Println(eachCons([]int{1, 2, 3, 4, 5, 6}, 3))
	// fmt.Println(reverseNum(0))
	// fmt.Println(keepHydrated(11.8))
	// fmt.Println(sortStar([]string{"turns", "out", "random", "test", "cases", "are", "easier", "than", "writing", "out", "basic", "ones"}))
	// fmt.Println(terminalGame(3, 6))
	// fmt.Println(leaveTheater(16, 11, 5, 3))
	// fmt.Println(areaVolume(4, 2, 6))
	// fmt.Println(nearestSquareNumber(10))
	// fmt.Println(removeStringSpaces("8 j 8   mBliB8g  imjB8B8  jl  B"))
	// fmt.Println(replaceDots("one.two.three"))
	// fmt.Println(binToDec("1001001"))
	// fmt.Println(centuryFromYear(2742))
	// fmt.Println(countOddBelowN(7))
	// fmt.Println(squareOrNot([]int{4, 3, 9, 7, 2, 1}))
	// fmt.Println(getMatrix(9))
	fmt.Println(fib(5))
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

/*
Some new animals have arrived at the zoo. The zoo keeper is concerned that perhaps the animals do not have the right tails.
To help her, you must correct the broken function to make sure that the second argument (tail), is the same as the last
letter of the first argument (body) - otherwise the tail wouldn't fit!

If the tail is right return true, else return false.

The arguments will always be non empty strings, and normal letters.

best practice

	func CorrectTail(body string, tail rune) bool {
		return rune(body[len(body)-1]) == tail
	}
*/
func zooTails(body string, tail rune) bool {
	if body == "" || tail == 0 {
		return false
	}

	if rune(body[len(body)-1]) != tail {
		return false
	}

	return true
}

/*
Create a method each_cons that accepts a list and a number n, and returns cascading subsets of the list of size n,
like so:

each_cons([1,2,3,4], 2)
  #=> [[1,2], [2,3], [3,4]]

each_cons([1,2,3,4], 3)
  #=> [[1,2,3], [2,3,4]]

As you can see, the lists are cascading; ie, they overlap, but never out of order.
*/

func eachCons(array []int, n int) [][]int {
	var res [][]int
	for i := 0; i <= len(array)-n; i++ {
		res = append(res, array[i:i+n])
	}
	return res
}

/*
Convert number to reversed array of digits
Given a random non-negative number, you have to return the digits of this number within an array in reverse order.

35231 => [1,3,2,5,3]
0 => [0]
*/
func reverseNum(number int) []int {
	if number == 0 {
		return []int{0}
	}

	result := make([]int, 0, len(strconv.Itoa(number)))
	for number > 0 {
		digit := number % 10
		number = number / 10
		result = append(result, digit)
	}
	return result
}

/*
Nathan loves cycling.

Because Nathan knows it is important to stay hydrated, he drinks 0.5 litres of water per hour of cycling.
You get given the time in hours and you need to return the number of litres Nathan will drink, rounded
to the smallest value.

For example:
time = 3 ----> litres = 1
time = 6.7---> litres = 3
time = 11.8--> litres = 5
*/
func keepHydrated(time float64) int {
	if time <= 0 {
		return 0
	}

	return int(time * 0.5)
}

/*
You will be given a list of strings. You must sort it alphabetically (case-sensitive, and based on the ASCII values of the chars)
and then return the first value.
The returned value must be a string, and have "***" between each of its letters.
You should not remove or add elements from/to the array.

best practice

	func TwoSort(arr []string) string {
		sort.Strings(arr)
		chars := strings.Split(arr[0], "")
		return strings.Join(chars, "***")
	}
*/
func sortStar(array []string) string {
	var result string
	if len(array) == 0 {
		return result
	} else {
		sort.Strings(array)
		count := 0
		for _, v := range array[0] {
			if count == 0 {
				s := string(v)
				result += s
			} else {
				result += "***" + string(v)
			}
			count++
		}
	}
	return result
}

/*
Terminal game move function
In this game, the hero moves from left to right. The player rolls the dice and moves the number of spaces
indicated by the dice two times.

Create a function for the terminal game that takes the current position of the hero and the roll (1-6)
and return the new position.

Example
move(3, 6) should equal 15
*/
func terminalGame(position, roll int) int {
	return roll*2 + position
}

/*
Your friend advised you to see a new performance in the most popular theater in the city.
He knows a lot about art and his advice is usually good, but not this time: the performance turned out
to be awfully dull. It's so bad you want to sneak out, which is quite simple, especially since the exit is
located right behind your row to the left. All you need to do is climb over your seat and make your way to
the exit.

The main problem is your shyness: you're afraid that you'll end up blocking the view (even if only for a couple
of seconds) of all the people who sit behind you and in your column or the columns to your left. To gain some
courage, you decide to calculate the number of such people and see if you can possibly make it to the exit
without disturbing too many people.

Given the total number of rows and columns in the theater (nRows and nCols, respectively), and the row and
column you're sitting in, return the number of people who sit strictly behind you and in your column or to the
left, assuming all seats are occupied.

For nCols = 16, nRows = 11, col = 5 and row = 3, the output should be 96.

Input/Output
[input] integer nCols
An integer, the number of theater's columns.
Constraints: 1 ≤ nCols ≤ 1000.

[input] integer nRows
An integer, the number of theater's rows.
Constraints: 1 ≤ nRows ≤ 1000.

[input] integer col
An integer, the column number of your own seat (with the rightmost column having index 1).
Constraints: 1 ≤ col ≤ nCols.

[input] integer row
An integer, the row number of your own seat (with the front row having index 1).
Constraints: 1 ≤ row ≤ nRows.

[output] an integer
The number of people who sit strictly behind you and in your column or to the left.
*/
func leaveTheater(nCols, nRows, col, row int) int {
	return (nCols - col + 1) * (nRows - row)
}

/*
Write a function that returns the total surface area and volume of a box as an array: [area, volume]
*/
func areaVolume(w, h, d int) [2]int {
	s := 2 * (w*h + w*d + h*d)
	v := w * h * d
	return [2]int{s, v}
}

/*
Your task is to find the nearest square number of a positive integer n. In mathematics, a square number or perfect square
is an integer that is the square of an integer; in other words, it is the product of some integer with itself.

For example, if n = 111, then the nearest square number equals 121, since 111 is closer to 121, the square of 11,
than 100, the square of 10.

If n is already a perfect square (e.g. n = 144, n = 81, etc.), you need to just return n.
*/
func nearestSquareNumber(n int) int {
	n = int(math.Round(math.Sqrt(float64(n))))
	return n * n
}

/*
Write a function that removes the spaces from the string, then return the resultant string.

Examples (Input -> Output):
"8 j 8   mBliB8g  imjB8B8  jl  B" -> "8j8mBliB8gimjB8B8jlB"
"8 8 Bi fk8h B 8 BB8B B B  B888 c hl8 BhB fd" -> "88Bifk8hB8BB8BBBB888chl8BhBfd"
"8aaaaa dddd r     " -> "8aaaaaddddr"
*/
func removeStringSpaces(str string) string {
	return strings.Replace(str, " ", "", -1)
}

/*
The code provided is supposed replace all the dots . in the specified String str with dashes -
But it's not working properly.

Task
Fix the bug so we can all go home early.

Notes
String str will never be null.
*/
func replaceDots(str string) string {
	return strings.ReplaceAll(str, ".", "-")
}

/*
Complete the function which converts a binary number (given as a string) to a decimal number.
Expect(BinToDec("10")).To(Equal(2))
Expect(BinToDec("11")).To(Equal(3))
Expect(BinToDec("101010")).To(Equal(42))
Expect(BinToDec("1001001")).To(Equal(73))
*/
func binToDec(bin string) int {
	res, _ := strconv.ParseInt(bin, 2, 64)
	return int(res)
}

/*
Century From Year

The first century spans from the year 1 up to and including the year 100,
the second century - from the year 101 up to and including the year 200, etc.
Given a year, return the century it is in.

1705 --> 18
1900 --> 19
1601 --> 17
2000 --> 20
2742 --> 28

best practice

	func century(year int) int {
		return (year + 99)/100
	}
*/
func centuryFromYear(year int) int {
	if year == 0 {
		return 1
	}

	if year%100 == 0 {
		return year / 100
	} else {
		return year/100 + 1
	}
}

/*
Count Odd Numbers below n
Given a number n, return the number of positive odd numbers below n, EASY!

7  -> 3 (because odd numbers below 7 are [1, 3, 5])
15 -> 7 (because odd numbers below 15 are [1, 3, 5, 7, 9, 11, 13])
Expect large Inputs!
*/
func countOddBelowN(n int) int {
	return n / 2
}

/*
Write a method, that will get an integer array as parameter and will process every number from this array.
Return a new array with processing every number of the input-array like this:
If the number has an integer square root, take this, otherwise square the number.

[4,3,9,7,2,1] -> [2,9,3,49,4,1]
The input array will always contain only positive numbers, and will never be empty or null.

best

	func SquareOrSquareRoot(arr []int) []int {
		ret := []int{}
		for _, n := range arr {
			i, f := math.Modf(math.Sqrt(float64(n)))
			if f == 0 {
				ret = append(ret, int(i))
			} else {
				ret = append(ret, n*n)
			}
		}
		return ret
	}
*/
func squareOrNot(arr []int) []int {
	result := make([]int, 0)
	for _, v := range arr {
		if num := math.Sqrt(float64(v)); num == float64(int(num)) {
			result = append(result, int(num))
		} else {
			result = append(result, v*v)
		}
	}
	return result
}

/*
// GetUniqInt64sStableFront Возвращает набор уникальных чисел, сохраняя порядок первых встреченных значений
func GetUniqInt64sStableFront(in []int64) []int64 {
	uniqMap := make(map[int64]struct{}, len(in))

	res := make([]int64, 0, len(in))

	for _, num := range in {
		_, ok := uniqMap[num]
		// Если элемент уже был обработан - пропускаем
		if ok {
			continue
		}

		uniqMap[num] = struct{}{}

		res = append(res, num)
	}

	return res
}
*/

/*
Пример работы с пакетом reflect
func convertToArray(row reportModel.Row) []string {
	v := reflect.ValueOf(row)
	t := reflect.TypeOf(row)
	values := make([]string, 0, v.NumField())
	hyperLink := false
	titleIndex := 0

	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Kind() == reflect.Slice {
			continue
		}
		if v.Field(i).Kind() != reflect.String {
			continue
		}
		// if has hyperlink, skip link
		if t.Field(i).Name == "HyperLink" && v.Field(i).String() != "" {
			hyperLink = true
			continue
		}
		fieldValue := fmt.Sprint(v.Field(i).Interface())
		if len(fieldValue) > 0 {
			values = append(values, fieldValue)
			if t.Field(i).Name == "Title" {
				titleIndex = len(values) - 1
			}
		}
	}

	// if has hyperlink, skip title
	if hyperLink {
		values = slices.Delete(values, titleIndex, titleIndex+1)
	}

	return values
}
*/

/*
Написать функцию, которая вернет матрицу 3х3, заполненную уникальными случайными числами от 1 до N.
Signature функции: func getMatrix() [][]int {}
Пример матрицы если N=20:
    [ [ 4,  7,  13 ],
      [ 14, 3,  6  ],
      [ 1,  18, 10 ] ]
*/

func getMatrix(n int) [][]int {
	if n < 9 {
		return [][]int{}
	}

	uniqueNumbers := make(map[int]struct{})
	var numbers []int

	for len(numbers) < n {
		num := rand.Intn(n) + 1

		if _, exists := uniqueNumbers[num]; !exists {
			uniqueNumbers[num] = struct{}{}
			numbers = append(numbers, num)
		}
	}

	matrix := make([][]int, 3)
	for i := 0; i < 3; i++ {
		matrix[i] = numbers[i*3 : (i+1)*3]
	}

	return matrix
}

// Fibonacci
func fib(n int) []int {
	res := make([]int, 0, n+1)
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
		res = append(res, y)
	}
	return res
}
