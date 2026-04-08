package profile

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoad_UnknownPreset(t *testing.T) {
	_, err := Load("ableton@12", "does-not-exist")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestLoad_SamplesIgnoredIgnoreGolden(t *testing.T) {
	rp, err := Load("ableton@12", "samples-ignored")
	if err != nil {
		t.Fatal(err)
	}
	golden := filepath.Join("testdata", "golden", "samples-ignored.ignore")
	want, err := os.ReadFile(golden)
	if err != nil {
		t.Fatal(err)
	}
	got := strings.TrimRight(rp.Ignore, "\n\r")
	exp := strings.TrimRight(string(want), "\n\r")
	if got != exp {
		t.Fatalf("ignore mismatch\n--- got ---\n%s\n--- want ---\n%s", rp.Ignore, want)
	}
}

func TestBundleDigestStablePrefix(t *testing.T) {
	d := BundleDigest()
	if len(d) < 8 || d[:7] != "sha256:" {
		t.Fatalf("unexpected digest: %q", d)
	}
}
