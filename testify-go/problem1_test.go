package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProblem1(t *testing.T) {
	actual := Problem1(1, 2)
	expexted := 3
	t.Run("case bilangan positif", func(t *testing.T) {
		assert.Equal(t, expexted, actual, "Error, Hasil tidak sesuai")
	})

	actualNegative := Problem1(-1, -2)
	expextedNegative := -3
	t.Run("case bilangan positif", func(t *testing.T) {
		assert.Equal(t, expextedNegative, actualNegative, "Error, Hasil tidak sesuai")
	})
}
