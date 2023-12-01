package main

import "testing"

func testLineValueExtractor(t *testing.T) {
	var line string = "d2six5dmlqczzrtp79brzzq"

	var got = ExtractValueFromLine(line)
	var want = 29

	if got != want {
		t.Errorf("got %q, wanted %q.", got, want)
	}
}
