package ast

import (
	"testing"

	"github.com/jeremi-traverse/monkey/token"
)

// Testing with single statement let myVar = anotherVar;
func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	expectedProgramString := "let myVar = anotherVar;"
	if program.String() != expectedProgramString {
		t.Errorf("program.String() wrong. expected %s got : %s",
			expectedProgramString,
			program.String())
	}
}
