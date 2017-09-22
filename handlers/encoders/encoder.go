package encoders

// Encoder is an interface to encode HTTP responses.
type Encoder interface {
	// Encode is used to encode HTTP responses.
	Encode(v interface{}) ([]byte, error)

	// ContentType returns Content-Type to represent the encode format.
	ContentType() string
}
