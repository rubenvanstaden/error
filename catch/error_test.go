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
	}

	for _, tc := range tests {
		act := tc.err.Error()
		test.Equals(t, tc.exp, act)
	}
}

func TestUnit_Debug(t *testing.T) {

	e0 := &catch.Debug{Op: "catch.t0", Err: fmt.Errorf("path not found")}
	e1 := &catch.Debug{Op: "catch.t1", Err: fmt.Errorf("internal error, %s", e0)}

	ee := &catch.Error{Code: catch.NOTFOUND, Message: "cannot find task"}
	e2 := &catch.Debug{Op: "catch.t1", Err: fmt.Errorf("internal error, %s", ee)}

	tests := []errorTester{
		{
			desc: "With Code, and message string",
			err:  e0,
			exp:  "catch.t0: path not found",
		},
		{
			desc: "With Code, and message string",
			err:  e1,
			exp:  "catch.t1: internal error, catch.t0: path not found",
		},
		{
			desc: "With Code, and message string",
			err:  e2,
			exp:  "catch.t1: internal error, NOT_FOUND: cannot find task",
		},
	}

	for _, tc := range tests {
		act := tc.err.Error()
		fmt.Println(act)
		test.Equals(t, tc.exp, act)
	}
}
