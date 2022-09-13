package chunk

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNew(t *testing.T) {
	r := strings.NewReader(`hoge
fuga
`)
	chunk, err := NewChunk(r)
	if err != nil {
		t.Fatal("failed creating a chunk:", err)
	}

	if diff := cmp.Diff([]byte("hoge"), chunk.Head()); diff != "" {
		t.Error("expected no diff but got:", diff)
	}

	if actual := chunk.Next(); !actual {
		t.Error("expected true but got", actual)
	}

	if diff := cmp.Diff([]byte("fuga"), chunk.Head()); diff != "" {
		t.Error("expected no diff but got:", diff)
	}

	if actual := chunk.Next(); actual {
		t.Error("expected false but got", actual)
	}

	if head := chunk.Head(); head != nil {
		t.Error("expected nil but got:", head)
	}
}
