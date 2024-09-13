package defeditor

import (
	"math/rand/v2"
	"testing"
)

func TestOpen(t *testing.T) {
	n := rand.Int()
	s, err := Open()
	if err != nil {
		t.Errorf("open_editor: %s\n", err)
	}
	t.Logf("%s\n", s)
}
