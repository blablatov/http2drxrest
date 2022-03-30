package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var tests = []struct {
		Name string
		Doc  string
	}{
		{"", "\n"},
		{".", "\t"},
		{"\t", "NaN\null\n\n"},
		{"Data for test", "Number 99999 to data test"},
		{"Yes, no", "No, or, yes, _, ops"},
	}

	var prevName string
	for _, test := range tests {
		if test.Name != prevName {
			fmt.Printf("\n%s\n", test.Name)
			prevName = test.Name
		}
	}

	var prevDoc string
	for _, test := range tests {
		if test.Doc != prevDoc {
			fmt.Printf("\n%s\n", test.Doc)
			prevDoc = test.Doc
		}
	}
}
