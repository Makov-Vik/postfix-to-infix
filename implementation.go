package lab2

import (
	"errors"
	"strings"
)

func PostfixToInfix(input string) (string, error) {
	inputArr := strings.Split(input, " ")

	var signCount int
	var numberCount int
	buffer := make([]string, 0)

	for _, operator := range inputArr {
		if operator == "^" || operator == "*" || operator == "/" || operator == "+" || operator == "-" {
			signCount++
		} else {
			numberCount++
		}
	}

	if signCount != numberCount-1 {
		return "", errors.New("Error: Invalid postfix equation")
	}

	for i := 0; i <= len(inputArr)-1; i++ {

		if inputArr[i] != "^" && inputArr[i] != "*" && inputArr[i] != "/" && inputArr[i] != "+" && inputArr[i] != "-" {
			buffer = append(buffer, inputArr[i])
		} else {
			if len(buffer) >= 2 {
				operand1 := buffer[len(buffer)-1]
				buffer = buffer[:len(buffer)-1]
				operand2 := buffer[len(buffer)-1]
				buffer = buffer[:len(buffer)-1]
				buffer = append(buffer, "("+operand2+inputArr[i]+operand1+")")
			} else {
				return "", errors.New("Error: Invalid postfix equation")
			}
		}
	}

	return buffer[0][1 : len(buffer[0])-1], nil
}
