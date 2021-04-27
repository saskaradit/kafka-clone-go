package client

import (
	"bytes"
	"testing"
)

func TestCutToLast(t *testing.T) {
	res := []byte("100\n101\n10")

	wantTruncated, wantRest := []byte("100\n101\n"), []byte("10")
	gotTruncated, gotRest, err := cutToLast(res)
	if err != nil {
		t.Errorf("cutToLast(%q): got error %v; want no errors", string(res), err)
	}

	if !bytes.Equal(gotTruncated, wantTruncated) || !bytes.Equal(gotRest, wantRest) {
		t.Errorf("cutToLast(%q): got %q, %q; want %q %q ", string(res), string(gotTruncated), string(gotRest), string(wantTruncated), string(wantRest))
	}

}
