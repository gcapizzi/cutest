package assert

type Result struct {
	Success bool
	Reason  string
}

type Matcher func(actual interface{}) Result

func That(actual interface{}, matcher Matcher) Result {
	return matcher(actual)
}

func (result Result) AndThat(actual interface{}, matcher Matcher) Result {
	if result.Success {
		return matcher(actual)
	}

	return result
}
