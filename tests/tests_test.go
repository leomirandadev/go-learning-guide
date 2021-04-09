package main

import (
	"errors"
	"testing"
)

func TestCalc(t *testing.T) {

	err := errors.New("B is bigger than A")

	type TestTableStruct struct {
		inputA         int
		inputB         int
		resultExpected int
		errExpected    error
	}

	testTable := []TestTableStruct{
		TestTableStruct{
			inputA:         2,
			inputB:         1,
			resultExpected: 1,
			errExpected:    nil,
		},
		TestTableStruct{
			inputA:         1,
			inputB:         1,
			resultExpected: 0,
			errExpected:    nil,
		},
		TestTableStruct{
			inputA:         0,
			inputB:         1,
			resultExpected: 0,
			errExpected:    err,
		},
	}

	for _, test := range testTable {
		result, err := calc(test.inputA, test.inputB)

		if test.resultExpected != result {
			t.Error("Test fail", result)
		}

		if test.errExpected != nil {
			if test.errExpected.Error() != err.Error() {
				t.Error("Test fail", err)
			}
		}
	}
}
