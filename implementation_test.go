package lab2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func SuccessfullyTestPostfixToInfix1(t *testing.T) {
	res, err := PostfixToInfix("234 2 * 3 + 2 /")
	if assert.Nil(t, err) {
		assert.Equal(t, "((234*2)+3)/2", res)
	}
}

func SuccessfullyTestPostfixToInfix2(t *testing.T) {
	res, err := PostfixToInfix("15 2 2.5 * / 32 + 2 /")
	if assert.Nil(t, err) {
		assert.Equal(t, "((15/(2*2.5))+32)/2", res)
	}
}

func SuccessfullyTestPostfixToInfix3(t *testing.T) {
	res, err := PostfixToInfix("32 1 + 15 3 / 2 * 1 / 22 * -")
	if assert.Nil(t, err) {
		assert.Equal(t, "(32+1)-((((15/3)*2)/1)*22)", res)
	}
}

func SuccessfullyTestPostfixToInfix4(t *testing.T) {
	res, err := PostfixToInfix("0.75 225 * 1 - 2 / 32 - 38 2 * +")
	if assert.Nil(t, err) {
		assert.Equal(t, "((((0.75*225)-1)/2)-32)+(38*2)", res)
	}
}

func SuccessfullyTestPostfixToInfix5(t *testing.T) {
	res, err := PostfixToInfix("1283 2 * 777 111 / * 228 114 / 113 1 + 2 / / +")
	if assert.Nil(t, err) {
		assert.Equal(t, "((1283*2)*(777/111))+((228/114)/((113+1)/2))", res)
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