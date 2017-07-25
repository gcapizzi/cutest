package test

import "testing"
import "github.com/gcapizzi/cutest/assert"

type Case func() assert.Result

type Suite struct {
	cases []Case
}

func NewSuite(cases []Case) Suite {
	return Suite{cases: cases}
}

func (suite *Suite) Run() assert.Result {
	return suite.cases[0]()
}

func HandleResult(t *testing.T, result assert.Result) {
	if !result.Success {
		t.Error(result.Reason)
	}
}
