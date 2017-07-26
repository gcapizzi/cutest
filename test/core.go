package test

import "fmt"
import "testing"
import "github.com/gcapizzi/cutest/assert"

type Case func()

func NewCase(body func()) Case {
	return Case(body)
}

func (c Case) Run() (result assert.Result) {
	defer func() {
		recoveredValue := recover()
		if recoveredValue != nil {
			result = assert.Result{Success: false, Reason: fmt.Sprintf("%v", recoveredValue)}
		}
	}()

	c()

	result = assert.Result{Success: true}
	return
}

type Suite struct {
	cases []Case
}

func NewSuite(cases []Case) Suite {
	return Suite{cases: cases}
}

func (s *Suite) Run() assert.Result {
	return s.cases[0].Run()
}

func HandleResult(t *testing.T, result assert.Result) {
	if !result.Success {
		t.Error(result.Reason)
	}
}
