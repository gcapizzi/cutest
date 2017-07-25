package main

import "testing"
import "github.com/gcapizzi/cutest/test"
import "github.com/gcapizzi/cutest/assert"
import . "github.com/gcapizzi/cutest/assert/matchers"

func TestHelloWorld(t *testing.T) {
	testCases := []test.Case{
		func() assert.Result {
			return assert.That(true, Is(true)).
				AndThat(1, Is(2))
		},
	}

	testSuite := test.NewSuite(testCases)
	testResult := testSuite.Run()
	test.HandleResult(t, testResult)
}
