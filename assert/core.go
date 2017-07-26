package assert

type Result struct {
	Success bool
	Reason  string
}

type Matcher func(actual interface{}) Result

func That(actual interface{}, matcher Matcher) {
	result := matcher(actual)

	if !result.Success {
		panic(result.Reason)
	}
}
