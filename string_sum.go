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

func operandsParser(numbers []string, errMsg string, operands *[]int) (output string, err error) {
	for _, elem := range numbers {
		op, e := strconv.Atoi(string(elem))
		if e != nil {
			err := fmt.Errorf(errMsg, e)
			return "", err
		}
		operands = append(operands, op)
	}
	return "", nil
}

func StringSum(input string) (output string, err error) {

	if len(input) == strings.Count(input, " ") {
		err = fmt.Errorf("%w", errorEmptyInput)
		return "", err
	}

	var errMsg string = "bad token: %w"
	var operands []int
	var sum int

	input = strings.ReplaceAll(input, " ", "")
	flag := 1

	if string(input[0]) == "-" {
		flag *= -1
	}

	input = strings.TrimLeft(input, "-")
	numPlus := strings.Split(input, "+")
	numMinus := strings.Split(input, "-")

	if len(numPlus) == 2 {
		output, err = operandsParser(numPlus, errMsg, &operands)
		if err != nil {
			return output, err
		}
		operands[0] *= flag
	} else if len(numMinus) == 2 {
		output, err = operandsParser(numMinus, errMsg, &operands)
		if err != nil {
			return output, err
		}
		operands[0] *= flag
		operands[1] *= -1
	} else {
		err = fmt.Errorf(errMsg, errorNotTwoOperands)
		return "", err
	}

	for _, operand := range operands {
		sum += operand
	}
	output = strconv.Itoa(sum)

	return output, nil
}
