package registry

import (
	"encoding/json"
	"testing"
)

// TestProbeEmptyListIsJSONArray asserts that listing an empty registry
// produces a JSON array ([]), not null. A nil result slice serializes to
// null, which breaks clients expecting an array.
func TestProbeEmptyListIsJSONArray(t *testing.T) {
	r := New()
	out := r.List()
	if out == nil {
		t.Fatalf("List() returned nil for empty registry; JSON would encode as null, want []")
	}
	b, err := json.Marshal(out)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if string(b) != "[]" {
		t.Fatalf("empty List JSON = %s, want []", b)
	}
}
