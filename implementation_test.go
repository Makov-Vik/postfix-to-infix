package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostfixToInfix(t *testing.T) {
	res, err := PostfixToInfix("17 10 + 3 * 9 /")
	if assert.Nil(t, err) {
		assert.Equal(t, "((17+10)*3)/9", res)
	}
}

func ExamplePostfixToInfix() {
	res, _ := PostfixToInfix("8 5 *")
	fmt.Println(res)

	// Output:
	// 8 5 *
}
