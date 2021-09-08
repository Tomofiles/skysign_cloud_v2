package json

import (
	"bytes"

	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/protobuf/runtime/protoiface"
)

var json_marshaler = jsonpb.Marshaler{
	EmitDefaults: true,
	OrigName:     true,
}

// Marshal .
func Marshal(m protoiface.MessageV1) []byte {
	var json bytes.Buffer
	json_marshaler.Marshal(&json, m)
	return json.Bytes()
}

// Unmarshal .
func Unmarshal(json []byte, m protoiface.MessageV1) error {
	return jsonpb.Unmarshal(bytes.NewReader(json), m)
}
