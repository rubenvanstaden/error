package catch_test

import (
	"fmt"
	"testing"

	"github.com/rubenvanstaden/error/catch"
	"github.com/rubenvanstaden/test"
)

type errorTester struct {
	desc string
	err  error
	exp  string
}

func TestUnit_Error(t *testing.T) {

	tests := []errorTester{
		{
			desc: "With Code, and message string",
			err:  &catch.Error{Code: catch.NOTFOUND, Message: "cannot find task"},
			exp:  "NOT_FOUND: cannot find task",
		},
		{
			desc: "With Code, and message string",
			err:  &catch.Error{Op: "catch.test", Err: fmt.Errorf("path not found")},
			exp:  "catch.test: path not found",
		},
	}

	for _, tc := range tests {
		act := tc.err.Error()
		test.Equals(t, tc.exp, act)
	}
}
