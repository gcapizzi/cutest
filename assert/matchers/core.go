package matchers

import "fmt"
import "github.com/gcapizzi/cutest/assert"

func Is(expected interface{}) assert.Matcher {
	return func(actual interface{}) assert.Result {
		if expected == actual {
			return assert.Result{Success: true}
		} else {
			return assert.Result{Success: false, Reason: fmt.Sprintf("%#v not equal to %#v", actual, expected)}
		}
	}
}
