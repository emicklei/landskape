package selfdiagnose

type recordingReporter struct {
	results []*Result
}

func (r *recordingReporter) Report(results []*Result) {
	r.results = results
}
