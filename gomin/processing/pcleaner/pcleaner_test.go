package pcleaner

import (
	"testing"
	"de.fraxxor.gofrax/gomin/processing/pfile"
)

func TestCleanFromImports_Identical(t *testing.T) {
	rows := []string{"first row", "second row"}
	sourcefile := pfile.Pfile{"", rows, "", ""}
	cleaner := PcleanerImpl{}
	cleanedfile := cleaner.CleanFromImports(&sourcefile)
	if !areRowsEqual((*cleanedfile).Rows, rows) {
		t.Errorf("Expected <%s> but was <%s>.\n", rows, (*cleanedfile).Rows)
	}
}

func TestCleanFromImports_Replaced(t *testing.T) {
	rows := []string{"import \"test\"", "test.doSmth"}
	expRows := []string{"doSmth"}
	sourcefile := pfile.Pfile{"", rows, "", ""}
	cleaner := PcleanerImpl{}
	cleanedfile := cleaner.CleanFromImports(&sourcefile)
	if !areRowsEqual((*cleanedfile).Rows, expRows) {
		t.Errorf("Expected <%s> but was <%s>.\n", expRows, (*cleanedfile).Rows)
	}
}

func areRowsEqual(exp []string, act []string) bool {
	if len(exp) != len(act) {
		return false
	}
	for i := range(exp) {
		if exp[i] != act[i] {
			return false
		}
	}
	return true
}