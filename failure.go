package eclair

type Failure struct {
	FailureType    string   `json:"failureType"`
	FailureMessage string   `json:"failureMessage"`
	FailedRoute    []string `json:"failedRoute"`
}
