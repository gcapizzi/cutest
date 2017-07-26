package main

import "testing"
import "github.com/gcapizzi/cutest/test"
import "github.com/gcapizzi/cutest/assert"
import . "github.com/gcapizzi/cutest/assert/matchers"

func TestHelloWorld(t *testing.T) {
	testCases := []test.Case{
		test.NewCase(func() {
			assert.That(true, Is(true))
			assert.That(42, Is(42))
		}),
	}

	testSuite := test.NewSuite(testCases)
	testResult := testSuite.Run()
	test.HandleResult(t, testResult)
}
