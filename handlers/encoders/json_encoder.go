package encoders

import "encoding/json"

// JSONEncoder is an encoder to encode HTTP responses in JSON format.
type JSONEncoder struct {
}

// Encode encodes given data in JSON format.
func (e JSONEncoder) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// ContentType returns Content-Type for JSON data.
func (e JSONEncoder) ContentType() string {
	return "application/json"
}
