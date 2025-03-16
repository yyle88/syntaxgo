package syntaxgo_astnorm

import "testing"

func TestStatementParts_MergeParts(t *testing.T) {
	a := StatementParts{
		"n int", "s string", "v float64",
	}
	t.Log(a.MergeParts())
}

func TestStatementLines_MergeLines(t *testing.T) {
	a := StatementLines{
		"var n int",
		"var s string",
		"var v float64",
	}
	t.Log(a.MergeLines())
}
