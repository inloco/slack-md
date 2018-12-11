package md

type Marshaler interface {
	MarshalSlackMD() ([]byte, error)
}
