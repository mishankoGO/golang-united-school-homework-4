package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use  these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {

	if len(input) == strings.Count(input, " ") {
		err := fmt.Errorf("bad token: %w", errorEmptyInput)
		return "", err
	}

	input = strings.ReplaceAll(input, " ", "")

	flag := 1

	if string(input[0]) == "-" {
		flag *= -1
	}

	inputTrimmed := strings.TrimLeft(input, "-")

	numPlus := strings.Split(inputTrimmed, "+")
	numMinus := strings.Split(inputTrimmed, "-")

	var operands []int

	if len(numPlus) == 2 {
		for _, elem := range numPlus {
			op, err := strconv.Atoi(string(elem))
			if err != nil {
				err = fmt.Errorf("bad token: %w", err)
				return "", err
			}
			operands = append(operands, op)
		}
		operands[0] *= flag
	} else if len(numMinus) == 2 {
		for _, elem := range numMinus {
			op, err := strconv.Atoi(elem)
			if err != nil {
				err = fmt.Errorf("bad token: %w", err)
				return "", err
			}
			operands = append(operands, op)
		}
		operands[0] *= flag
		operands[1] *= -1
	} else {
		err = fmt.Errorf("bad token: %w", errorNotTwoOperands)
		return "", err
	}
	var sum int
	for _, operand := range operands {
		sum += operand
	}
	return strconv.Itoa(sum), nil
}
