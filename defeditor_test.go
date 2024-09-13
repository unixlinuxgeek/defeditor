package defeditor

import (
	"math/rand/v2"
	"strconv"
	"testing"
)

func TestOpen(t *testing.T) {
	n := rand.Int()
	s, err := Open("tmp_" + strconv.Itoa(n) + ".txt")
	if err != nil {
		t.Errorf("open_editor: %s\n", err)
	}
	t.Logf("%s\n", s)
}
