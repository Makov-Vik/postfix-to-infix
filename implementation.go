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

	operArr := make([]string, signCount)
	var indxForSigns int = -1
	for i := 0; i <= len(inputArr)-1; i++ {
		if inputArr[i] == "^" || inputArr[i] == "*" || inputArr[i] == "/" || inputArr[i] == "+" || inputArr[i] == "-" {
			indxForSigns++
			operArr[indxForSigns] = inputArr[i]
		}
	}
	indxForSigns = 0

	for i := 0; i <= len(inputArr)-1; i++ {

		if inputArr[i] != "^" && inputArr[i] != "*" && inputArr[i] != "/" && inputArr[i] != "+" && inputArr[i] != "-" {
			buffer = append(buffer, inputArr[i])
		} else {
			if len(buffer) >= 2 {
				operand1 := buffer[len(buffer)-1]
				buffer = buffer[:len(buffer)-1]
				operand2 := buffer[len(buffer)-1]
				buffer = buffer[:len(buffer)-1]
				if operArr[indxForSigns] == "*" || operArr[indxForSigns] == "/" || operArr[indxForSigns] == "^" {
					buffer = append(buffer, ""+operand2+" "+operArr[indxForSigns]+" "+operand1+"")
				} else {
					if i == len(inputArr)-1 {
						buffer = append(buffer, ""+operand2+" "+operArr[indxForSigns]+" "+operand1+"")
					} else {
						if operArr[indxForSigns+1] == "*" || operArr[indxForSigns+1] == "/" || operArr[indxForSigns+1] == "^" {
							buffer = append(buffer, "("+operand2+" "+operArr[indxForSigns]+" "+operand1+")")
						} else {
							buffer = append(buffer, ""+operand2+" "+operArr[indxForSigns]+" "+operand1+"")
						}
					}
				}
				indxForSigns++
			} else {
				return "", errors.New("Error: Invalid postfix equation")
			}
		}
	}

	return buffer[0], nil
}
