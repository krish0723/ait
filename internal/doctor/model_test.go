package doctor

import "testing"

func TestSortFindings_order(t *testing.T) {
	fs := []Finding{
		{Code: "b", Severity: SeverityInfo, Path: "z"},
		{Code: "a", Severity: SeverityError, Path: "y"},
		{Code: "c", Severity: SeverityWarn, Path: "a"},
	}
	SortFindings(fs)
	if fs[0].Code != "a" || fs[0].Severity != SeverityError {
		t.Fatalf("want error first, got %+v", fs[0])
	}
	if fs[1].Severity != SeverityWarn {
		t.Fatalf("want warn second: %+v", fs[1])
	}
	if fs[2].Severity != SeverityInfo {
		t.Fatalf("want info last: %+v", fs[2])
	}
}
