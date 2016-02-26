package pqarrays

import (
	"testing"
)

var testInputs = map[string][]token{
	``:                                       []token{{typ: tokenError, val: "expected array to start before "}},
	`{}`:                                     []token{{typ: tokenArrayStart, val: "{"}, {typ: tokenArrayEnd, val: "}"}},
	`{    }`:                                 []token{{typ: tokenArrayStart, val: "{"}, {typ: tokenWhitespace, val: "    "}, {typ: tokenArrayEnd, val: "}"}},
	`{lions}`:                                []token{{typ: tokenArrayStart, val: "{"}, {typ: tokenString, val: "lions"}, {typ: tokenArrayEnd, val: "}"}},
	`{lions,tigers}`:                         []token{{typ: tokenArrayStart, val: "{"}, {typ: tokenString, val: "lions"}, {typ: tokenSeparator, val: ","}, {typ: tokenString, val: "tigers"}, {typ: tokenArrayEnd, val: "}"}},
	`{lions,tigers,bears}`:                   []token{{typ: tokenArrayStart, val: "{"}, {typ: tokenString, val: "lions"}, {typ: tokenSeparator, val: ","}, {typ: tokenString, val: "tigers"}, {typ: tokenSeparator, val: ","}, {typ: tokenString, val: "bears"}, {typ: tokenArrayEnd, val: "}"}},
	`{lions,tigers,bears,"oh my!"}`:          []token{{typ: tokenArrayStart, val: "{"}, {typ: tokenString, val: "lions"}, {typ: tokenSeparator, val: ","}, {typ: tokenString, val: "tigers"}, {typ: tokenSeparator, val: ","}, {typ: tokenString, val: "bears"}, {typ: tokenSeparator, val: ","}, {typ: tokenString, val: "oh my!"}, {typ: tokenArrayEnd, val: "}"}},
	`{{two,dimensional},{array,"of items"}}`: []token{{typ: tokenArrayStart, val: "{"}, {typ: tokenArrayStart, val: "{"}, {typ: tokenString, val: "two"}, {typ: tokenSeparator, val: ","}, {typ: tokenString, val: "dimensional"}, {typ: tokenArrayEnd, val: "}"}, {typ: tokenSeparator, val: ","}, {typ: tokenArrayStart, val: "{"}, {typ: tokenString, val: "array"}, {typ: tokenSeparator, val: ","}, {typ: tokenString, val: "of items"}, {typ: tokenArrayEnd, val: "}"}, {typ: tokenArrayEnd, val: "}"}},
}

func TestInputsTable(t *testing.T) {
	for input, expectedTokens := range testInputs {
		l := lex(input)
		var tokens []token
		for {
			tok := l.nextToken()
			if tok.typ == tokenEOF {
				break
			}
			tokens = append(tokens, tok)
			if tok.typ == tokenError {
				break
			}
		}
		t.Logf("`%s`: %#+v\n", input, tokens)
		if len(tokens) != len(expectedTokens) {
			t.Fatalf("Expected %d tokens, got %d\n", len(expectedTokens), len(tokens))
		}
		for pos, tok := range tokens {
			if expectedTokens[pos].typ != tok.typ {
				t.Errorf("Expected token in pos %d to have type of %s, got %s instead.", pos, expectedTokens[pos].typ, tok.typ)
			}
			if expectedTokens[pos].val != tok.val {
				t.Errorf("Expected token in pos %d to have value of `%s`, got `%s` instead.", pos, expectedTokens[pos].val, tok.val)
			}
		}
	}
}
