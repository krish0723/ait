package aitinit

import (
	"bytes"
	"testing"
)

func TestMergeIntoFile_create(t *testing.T) {
	out, err := MergeIntoFile(nil, "foo\n", false)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(out, []byte("foo")) || !bytes.Contains(out, []byte(markerBegin)) {
		t.Fatalf("%s", out)
	}
}

func TestMergeIntoFile_append(t *testing.T) {
	existing := []byte("manual\n")
	out, err := MergeIntoFile(existing, "bar\n", false)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.HasPrefix(out, []byte("manual\n\n")) {
		t.Fatalf("%q", out)
	}
}

func TestMergeIntoFile_replace(t *testing.T) {
	existing := []byte("top\n" + AitBlock("old") + "tail\n")
	out, err := MergeIntoFile(existing, "new\n", false)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Contains(out, []byte("old")) {
		t.Fatalf("should replace inner: %s", out)
	}
	if !bytes.Contains(out, []byte("new")) || !bytes.Contains(out, []byte("top")) || !bytes.Contains(out, []byte("tail")) {
		t.Fatalf("%s", out)
	}
}

func TestMergeIntoFile_idempotent(t *testing.T) {
	body := "x\n"
	first, err := MergeIntoFile(nil, body, false)
	if err != nil {
		t.Fatal(err)
	}
	second, err := MergeIntoFile(first, body, false)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(first, second) {
		t.Fatalf("first:\n%s\nsecond:\n%s", first, second)
	}
}

func TestMergeIntoFile_duplicateMarkers(t *testing.T) {
	bad := []byte(AitBlock("a") + "\n" + "# BEGIN ait\n")
	_, err := MergeIntoFile(bad, "b", false)
	if err != ErrMergeConflict {
		t.Fatalf("want conflict, got %v", err)
	}
}

func TestMergeIntoFile_forceDup(t *testing.T) {
	bad := []byte(AitBlock("a") + "\n" + markerBegin + "\norphan\n")
	out, err := MergeIntoFile(bad, "b", true)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Count(out, []byte(markerBegin)) != 1 {
		t.Fatalf("expected single begin: %s", out)
	}
	if !bytes.Contains(out, []byte("b")) {
		t.Fatal(out)
	}
}
