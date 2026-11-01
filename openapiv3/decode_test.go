package openapi_v3

import "testing"

// TestDecodeOwnsStrings checks that decoded documents survive reuse of the wire buffer.
func TestDecodeOwnsStrings(t *testing.T) {
	// Encode nested strings so both document and child message decoders run.
	original := &Document{Openapi: "3.0.0", Info: &Info{Title: "example", Version: "v1"}}
	data, err := original.MarshalVT()
	if err != nil {
		t.Fatal(err)
	}

	// Decode through the legacy entry point and then overwrite its input buffer.
	var decoded Document
	if err := decoded.UnmarshalVTUnsafe(data); err != nil {
		t.Fatal(err)
	}
	clear(data)

	// The decoded document owns every string independently of that buffer.
	if !original.EqualVT(&decoded) {
		t.Fatalf("decoded document changed after input reuse: %v", &decoded)
	}
}
