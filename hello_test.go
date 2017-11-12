package gopathless

import (
	"bytes"
	"testing"
)

func TestHelloLibrary(t *testing.T) {
	var buf bytes.Buffer
	HelloLibrary(&buf)
	if buf.String() == "heji, mundus!\n" {
		t.Log("test passed!")
	} else {
		t.Fatalf("test failed :(")
	}
}
