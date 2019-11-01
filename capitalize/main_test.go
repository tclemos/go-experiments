package main

import (
	"encoding/json"
	"testing"
)

func TestCapitalize(T *testing.T) {

	type testCase struct {
		name     string
		input    string
		expected string
	}

	testCases := []testCase{
		testCase{
			name:     "simple",
			input:    `{"p1":"this value should be capitalized"}`,
			expected: `{"p1":"This Value Should Be Capitalized"}`,
		},
		testCase{
			name:     "complex",
			input:    `{"o1":{"p1":"this value should be capitalized","p2":"this value should be capitalized"},"p1":"this value should be capitalized"}`,
			expected: `{"o1":{"p1":"This Value Should Be Capitalized","p2":"This Value Should Be Capitalized"},"p1":"This Value Should Be Capitalized"}`,
		},
		testCase{
			name:     "multilevel",
			input:    `{"o1":{"o2":{"p1":"this value should be capitalized","p2":"this value should be capitalized"},"p1":"this value should be capitalized","p2":"this value should be capitalized"},"p1":"this value should be capitalized"}`,
			expected: `{"o1":{"o2":{"p1":"This Value Should Be Capitalized","p2":"This Value Should Be Capitalized"},"p1":"This Value Should Be Capitalized","p2":"This Value Should Be Capitalized"},"p1":"This Value Should Be Capitalized"}`,
		},
	}

	for _, testCase := range testCases {
		T.Logf("initializing test %s\n", testCase.name)
		m := make(map[string]interface{}, 0)
		json.Unmarshal([]byte(testCase.input), &m)
		capitalizeMap(m)

		r, _ := json.Marshal(m)

		if string(r) != testCase.expected {
			T.Errorf("Test case %s failed.\nresult:\n%s\nexpected:\n%s", testCase.name, r, testCase.expected)
		}
	}

}
