package lab2

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	input := strings.NewReader("")
	output := bytes.NewBufferString("")

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	if assert.Nil(t, err) {
		assert.Equal(t, "", output.String())
	}
}

func Test2(t *testing.T) {
	input := strings.NewReader("15 2 7 * / 32 + 2 /")
	output := bytes.NewBufferString("")

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	if assert.Nil(t, err) {
		assert.Equal(t, "(15 / 2 * 7 + 32) / 2", output.String())
	}
}

func Test3(t *testing.T) {
	input := strings.NewReader("32 1 15 - 3 / 2 * 1 / 22 * +")
	output := bytes.NewBufferString("")

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	if assert.Nil(t, err) {
		assert.Equal(t, "32 + (1 - 15) / 3 * 2 / 1 * 22", output.String())
	}
}

func Test4(t *testing.T) {
	input := strings.NewReader("0.75 225 * 1 - 2 / 32 - 38 2 * +")
	output := bytes.NewBufferString("")

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	if assert.Nil(t, err) {
		assert.Equal(t, "((0.75 * 225 - 1) / 2 - 32) + 38 * 2", output.String())
	}
}


func Test5(t *testing.T) {
	input := strings.NewReader(" ")
	output := bytes.NewBufferString("")

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	assert.NotNil(t, err)
}

func Test6(t *testing.T) {
	input := strings.NewReader("+")
	output := bytes.NewBufferString("")

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	assert.NotNil(t, err)
}

func Test7(t *testing.T) {
	input := strings.NewReader("8 +")
	output := bytes.NewBufferString("")

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	assert.NotNil(t, err)
}

func Test8(t *testing.T) {
	input := strings.NewReader("0.75")
	output := bytes.NewBufferString("")

	handler := &ComputeHandler{
		Input:  input,
		Output: output,
	}

	err := handler.Compute()
	if assert.Nil(t, err) {
		assert.Equal(t, "0.75", output.String())
	}
}