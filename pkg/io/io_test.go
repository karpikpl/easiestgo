package io

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExecuteActionOnEachLine_Should_CallPassedActionWManyTimes(t *testing.T) {
	// Arrange
	testData := "3029\n4\n5\n42\n0"
	var calledTimes int
	dummyAction := func(n int) {
		calledTimes++
		t.Log("read n=", n)
	}

	// Act
	ExecuteActionOnEachLine(strings.NewReader(testData), dummyAction)

	// Assert
	assert.Equal(t, 4, calledTimes)
}
