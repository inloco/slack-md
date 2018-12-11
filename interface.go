package slack_md

type Marshaler interface {
	MarshalSlackMD() ([]byte, error)
}
