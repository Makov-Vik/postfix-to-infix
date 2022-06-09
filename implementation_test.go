package lab2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func SuccessfullyTestPostfixToInfix1(t *testing.T) {
	res, err := PostfixToInfix("2 3 *")
	if assert.Nil(t, err) {
		assert.Equal(t, "2 * 3", res)
	}
}

func SuccessfullyTestPostfixToInfix2(t *testing.T) {
	res, err := PostfixToInfix("15 2 7 * / 32 + 2 /")
	if assert.Nil(t, err) {
		assert.Equal(t, "(15 / 2 * 7 + 32) / 2", res)
	}
}

func SuccessfullyTestPostfixToInfix3(t *testing.T) {
	res, err := PostfixToInfix("32 1 15 - 3 / 2 * 1 / 22 * +")
	if assert.Nil(t, err) {
		assert.Equal(t, "32 + (1 - 15) / 3 * 2 / 1 * 22", res)
	}
}

func SuccessfullyTestPostfixToInfix4(t *testing.T) {
	res, err := PostfixToInfix("0.75 225 * 1 - 2 / 32 - 38 2 * +")
	if assert.Nil(t, err) {
		assert.Equal(t, "((0.75 * 225 - 1) / 2 - 32) + 38 * 2", res)
	}
}

func SuccessfullyTestPostfixToInfix5(t *testing.T) {
	res, err := PostfixToInfix("3 10 + 7 * 12 5 / - 6 7 ^ * 142 2 / 35 - 5 ^ + 87 + 15 * 123 4 / +")
	if assert.Nil(t, err) {
		assert.Equal(t, "(((3 + 10) * 7 - 12 / 5) * 6 ^ 7 + (142 / 2 - 35) ^ 5 + 87) * 15 + 123 / 4", res)
	}
}

func ErrorTestPostfixToInfix1(t *testing.T) {
	_, err := PostfixToInfix("* 15 + 23 17 / + 53 62 + +")
	assert.NotNil(t, err)
}

func ErrorTestPostfixToInfix2(t *testing.T) {
	_, err := PostfixToInfix("1 2 3 4 5")
	assert.NotNil(t, err)
}

func ErrorTestPostfixToInfix3(t *testing.T) {
	_, err := PostfixToInfix(" ")
	assert.NotNil(t, err)
}

func ErrorTestPostfixToInfix4(t *testing.T) {
	_, err := PostfixToInfix("")
	assert.NotNil(t, err)
}
