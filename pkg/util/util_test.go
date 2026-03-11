package util

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct{
		input string
		expected []string
	}{
		{
			input : "hello world",
			expected: []string{"hello", "world",},
		},
		{
			input : "HELLO WORLD",
			expected: []string{"hello", "world",},
		},

	}

	for _, cs := range cases {
		res := CleanInput(cs.input)

		// check the lenght of the result
		if len(res) != len(cs.expected) {
			t.Errorf("The lenght are not same. Expected: %v, Actual: %v", len(cs.expected), len(res))
			continue
		}

		// check the content of the result
		for i := range res {
			resWrl := res[i]
			expWrl := cs.expected[i]
			if resWrl != expWrl {
				t.Errorf("Wrong Output. Expected: %v, Actual: %v", expWrl, resWrl)
			}
		}


	}
}
