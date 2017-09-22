package encoders

// Encoder is an interface to encode HTTP responses.
type Encoder interface {
	Encode(v interface{}) ([]byte, error)
}
