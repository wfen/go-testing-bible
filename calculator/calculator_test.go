package calculator_test

import (
	"testing"

	"github.com/wfen/go-testing-bible/calculator"
)

type TestCase struct {
	name     string
	value    int
	expected bool
	actual   bool
}

func TestCalculateIsArmstrong(t *testing.T) {
	t.Run("test for all 3 digit armstrong numbers", func(t *testing.T) {
		tests := []TestCase{
			TestCase{name: "Testing value for: 153", value: 153, expected: true},
			TestCase{name: "Testing value for: 370", value: 370, expected: true},
			TestCase{name: "Testing value for: 371", value: 371, expected: true},
			TestCase{name: "Testing value for: 407", value: 407, expected: true},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				actual := calculator.CalculateIsArmstrong(test.value)
				if test.expected != actual {
					t.Fail()
				}
			})
		}
	})
}

func TestNegativeCalculateIsArmstrong(t *testing.T) {
	t.Run("should return false for 350", func(t *testing.T) {
		testCase := TestCase{
			value:    350,
			expected: false,
		}
		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})

	t.Run("should return false for 300", func(t *testing.T) {
		testCase := TestCase{
			value:    300,
			expected: false,
		}
		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
		if testCase.actual != testCase.expected {
			t.Fail()
		}
	})
}
