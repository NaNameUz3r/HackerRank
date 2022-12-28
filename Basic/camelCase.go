// Camel Case is a naming style common in many programming languages. In Java, method and variable names typically start with a lowercase letter, with all subsequent words starting with a capital letter (example: startThread). Names of classes follow the same pattern, except that they start with a capital letter (example: BlueCar).

// Your task is to write a program that creates or splits Camel Case variable, method, and class names.

// Input Format

//     Each line of the input file will begin with an operation (S or C) followed by a semi-colon followed by M, C, or V followed by a semi-colon followed by the words you'll need to operate on.
//     The operation will either be S (split) or C (combine)
//     M indicates method, C indicates class, and V indicates variable
//     In the case of a split operation, the words will be a camel case method, class or variable name that you need to split into a space-delimited list of words starting with a lowercase letter.
//     In the case of a combine operation, the words will be a space-delimited list of words starting with lowercase letters that you need to combine into the appropriate camel case String. Methods should end with an empty set of parentheses to differentiate them from variable names.

// Output Format

//     For each input line, your program should print either the space-delimited list of words (in the case of a split operation) or the appropriate camel case string (in the case of a combine operation).

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputStr := scanner.Text()

		parametersRaw := inputStr[:4]
		parameters := strings.Split(parametersRaw, ";")
		strToProcess := inputStr[4:]

		if parameters[0] == "S" {
			splitString(strToProcess, parameters[1])
		} else if parameters[0] == "C" {
			combineString(strToProcess, parameters[1])
		}
	}

}

func splitString(str, entity string) {
	processStr := str
	var resultString string

	if entity == "M" {
		processStr = strings.TrimSuffix(processStr, "()")
	}

	i := 0
	for s := processStr; s != ""; s = s[i:] {
		i = strings.IndexFunc(s[1:], unicode.IsUpper) + 1
		if i <= 0 {
			i = len(s)
		}
		resultString += strings.ToLower(s[:i]) + " "
	}

	fmt.Println(resultString)
}

func combineString(str, entity string) {
	var resultStr string
	slice := strings.Split(str, " ")

	var idx int
	if entity == "C" {
		idx = 0
	} else {
		resultStr += slice[0]
		idx = 1
	}

	for i := idx; i < len(slice); i++ {
		runeScript := []rune(slice[i])
		runeScript[0] = unicode.ToUpper(runeScript[0])
		resultStr += string(runeScript)
	}

	if entity == "M" {
		resultStr += "()"
	}
	fmt.Println(resultStr)
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
