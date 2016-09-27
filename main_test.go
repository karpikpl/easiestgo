package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SumDigits_Should_SumDigits_When_1234(t *testing.T) {
	// Arrange
	testData := 1234

	// Act
	result := SumDigits(testData)

	// Assert
	assert.Equal(t, 10, result)
}

func Test_SumDigits_Should_SumDigits_When_5(t *testing.T) {
	// Arrange
	testData := 5

	// Act
	result := SumDigits(testData)

	// Assert
	assert.Equal(t, 5, result)
}

func Test_SumDigits_Should_SumDigits_When_11(t *testing.T) {
	// Arrange
	testData := 11

	// Act
	result := SumDigits(testData)

	// Assert
	assert.Equal(t, 2, result)
}
