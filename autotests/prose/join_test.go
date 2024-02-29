package prose

import "testing"

type testData struct {
	list     []string
	expected string
}

func TestOneTwoAndThreeElements(t *testing.T) {
	tests := []testData{
		{list: []string{}, expected: ""},
		{list: []string{"apples"}, expected: "apples"},
		{list: []string{"apples", "oranges"}, expected: "apples and oranges"},
		{list: []string{"apples", "oranges", "pears"}, expected: "apples, oranges, and pears"},
	}

	for _, test := range tests {
		executeTest(test.list, test.expected, t)
	}
}

func executeTest(list []string, expected string, t *testing.T) {
	actual := JoinWithCommas(list)

	passed := (actual == expected)

	if !passed {
		t.Error("Actual:", actual, "not as expected:", expected)
	}
}
