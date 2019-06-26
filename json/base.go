package json

import (
	"io"

	jsoniter "github.com/json-iterator/go"
)

var (
	j = jsoniter.ConfigCompatibleWithStandardLibrary

	Marshal    func(interface{}) ([]byte, error)
	Unmarshal  func([]byte, interface{}) error
	NewDecoder func(r io.Reader) *jsoniter.Decoder
	NewEncoder func(w io.Writer) *jsoniter.Encoder
)

func init() {
	Marshal = j.Marshal
	Unmarshal = j.Unmarshal
	NewDecoder = j.NewDecoder
	NewEncoder = j.NewEncoder
}
