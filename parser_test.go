package pqarrays

import (
	"testing"
)

func strPtr(in string) *string {
	return &in
}

var parseTestInputs = map[string][]*string{
	`{}`:                            []*string{},
	`{lions}`:                       []*string{strPtr("lions")},
	`{lions,tigers}`:                []*string{strPtr("lions"), strPtr("tigers")},
	`{lions,tigers,NULL}`:           []*string{strPtr("lions"), strPtr("tigers"), nil},
	`{lions,tigers,bears}`:          []*string{strPtr("lions"), strPtr("tigers"), strPtr("bears")},
	`{lions,tigers,bears,"oh my!"}`: []*string{strPtr("lions"), strPtr("tigers"), strPtr("bears"), strPtr("oh my!")},
}

func TestParseInputsTable(t *testing.T) {
	for input, expected := range parseTestInputs {
		l := lex(input)
		output, err := parse(l)
		if err != nil {
			t.Fatalf(err.Error())
		}
		t.Logf("`%s`: %#+v\n", input, output)
		if len(output) != len(expected) {
			t.Fatalf("Expected %d items in array, got %d\n", len(expected), len(output))
		}
		for pos, item := range output {
			if item == nil && expected[pos] != nil {
				t.Errorf("Expected %d to be %s, got nil instead.", pos, *expected[pos])
			} else if item != nil && expected[pos] == nil {
				t.Errorf("Expected %d to be nil, got %s instead.", pos, *item)
			} else if item != nil && expected[pos] != nil {
				continue
			} else if item == nil && expected[pos] == nil {
				continue
			} else if *item != *expected[pos] {
				t.Errorf("Expected %d to be %s, got %s instead.", pos, *expected[pos], *item)
			}
		}
	}
}
