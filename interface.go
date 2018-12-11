package md

// Marshaler is the interface implemented by types that
// can marshal themselves into custom Slack Markdown language.
type Marshaler interface {
	MarshalSlackMD() ([]byte, error)
}
