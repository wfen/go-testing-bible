package testy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterUnique(t *testing.T) {
	input := []Developer{
		Developer{Name: "Elliot"},
		Developer{Name: "Elliot"},
		Developer{Name: "David"},
		Developer{Name: "Alexander"},
		Developer{Name: "Eva"},
		Developer{Name: "Alan"},
	}

	expected := []string{
		"Elliot",
		"David",
		"Alexander",
		"Eva",
		"Alan",
	}

	result := FilterUnique(input)
	// ElementsMatch asserts that the specified listA is equal to specified listB ignoring the order of the elements.
	// If there are duplicate elements, the number of appearances of each of them in both lists should match
	assert.ElementsMatch(t, expected, result)
}

func TestNotFilterUnique(t *testing.T) {
	input := []Developer{
		Developer{Name: "Elliot"},
		Developer{Name: "Elliot"},
		Developer{Name: "David"},
		Developer{Name: "Alexander"},
		Developer{Name: "Eva"},
		Developer{Name: "Alan"},
	}

	expected := []string{
		"Elliot",
		"Eva",
		"Alan",
	}

	result := FilterUnique(input)
	assert.NotEqual(t, expected, result)
}
