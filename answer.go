package asker

// Answer is what is returned in response to a question when asked it.
type Answer struct {
	StringResponse string
	BoolResponse   bool
	Provided       bool
}
