package main

import (
	"strings"
	"testing"
)

func TestSort(t *testing.T) {
	tests := []struct {
		input, output string
	}{
		{
			// Basic
			input:  "",
			output: "",
		},
		{
			// Another basic test
			input:  "1.0.0",
			output: strings.Join([]string{"1.0.0", ""}, "\n"),
		},
		{
			// 0.1.1 is smaller than 1.0.0
			input:  "1.0.0\n0.1.1",
			output: strings.Join([]string{"0.1.1", "1.0.0", ""}, "\n"),
		},
		{
			// Un-jumble a bunch of semvers
			input: strings.Join([]string{
				"0.0.0",
				"1.4.0",
				"1.3.0",
				"1.2.0",
				"1.1.1",
			}, " "),
			output: strings.Join([]string{
				"0.0.0",
				"1.1.1",
				"1.2.0",
				"1.3.0",
				"1.4.0",
				""}, "\n"),
		},
		{
			// Release candidates are smaller semver numbers than the
			// release itself.
			input: strings.Join([]string{
				"1.2.3",
				"1.2.3-rc.1",
				"1.2.3-rc.0",
			}, " "),
			output: strings.Join([]string{
				"1.2.3-rc.0",
				"1.2.3-rc.1",
				"1.2.3",
				""}, "\n"),
		},
		{
			// Same as above, but we also accept a leading 'v'.  Doesn't affect
			// sorting, but does get printed the way it was specified.
			input: strings.Join([]string{
				"v1.2.3",
				"v1.2.3-rc.1",
				"1.2.3-rc.0",
			}, " "),
			output: strings.Join([]string{
				"1.2.3-rc.0",
				"v1.2.3-rc.1",
				"v1.2.3",
				""}, "\n"),
		},
		{
			// Same as above, but also includes nonsense strings that can't
			// be parsed as semver.
			input: strings.Join([]string{
				"v1.2.3",
				"vinnie",
				"v1.2.3-rc.1",
				"the",
				"1.2.3-rc.0",
				"pooh",
			}, " "),
			output: strings.Join([]string{
				"1.2.3-rc.0",
				"v1.2.3-rc.1",
				"v1.2.3",
				""}, "\n"),
		},
	}
	for _, test := range tests {
		var b strings.Builder
		err := SortSemver(strings.NewReader(test.input), &b)
		if err != nil {
			t.Errorf("for input: %v: unexpected error: %v", test.input, err)
		}
		if b.String() != test.output {
			t.Errorf("output: %q, want: %q", b.String(), test.output)
		}
	}
}
