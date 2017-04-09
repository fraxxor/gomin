package pcleaner

import (
	"testing"
	"de.fraxxor.gofrax/gomin/processing/pfile"
)

func TestClean_NothingToReplace(t *testing.T) {
	rows := []string{"first row", "second row"}
	file := pfile.Pfile{"", rows, []pfile.Goimport{}, "", ""}
	cleaner := CreateImportCleaner(&[]pfile.Pfile{file})
	cleaner.Clean(&file)
	if !areRowsEqual(file.Rows, rows) {
		t.Errorf("Expected <%s> but was <%s>.\n", rows, file.Rows)
	}
}

func TestClean_NoReplaceDueToDot(t *testing.T) {
	rows := []string{"doSmth"}
	expRows := []string{"doSmth"}
	refFile := pfile.Pfile{"", []string{}, []pfile.Goimport{}, "test", "test"}
	file := pfile.Pfile{"", rows, []pfile.Goimport{pfile.CreateGoimport("", "test")}, "", ""}
	cleaner := CreateImportCleaner(&[]pfile.Pfile{refFile, file})
	cleaner.Clean(&file)
	if !areRowsEqual(file.Rows, expRows) {
		t.Errorf("Expected <%s> but was <%s>.\n", expRows, file.Rows)
	}
}

func TestClean_NoReplaceDueToNoOwnPFile(t *testing.T) {
	rows := []string{"api.doSmth"}
	expRows := []string{"api.doSmth"}
	file := pfile.Pfile{"", rows, []pfile.Goimport{pfile.CreateGoimport("api", "api")}, "", ""}
	cleaner := CreateImportCleaner(&[]pfile.Pfile{file})
	cleaner.Clean(&file)
	if !areRowsEqual(file.Rows, expRows) {
		t.Errorf("Expected <%s> but was <%s>.\n", expRows, file.Rows)
	}
}

func TestClean_ReplaceSimple(t *testing.T) {
	rows := []string{"test.doSmth"}
	expRows := []string{"doSmth"}
	refFile := pfile.Pfile{"", []string{}, []pfile.Goimport{}, "test", "test"}
	file := pfile.Pfile{"", rows, []pfile.Goimport{pfile.CreateGoimport("test", "test")}, "", ""}
	cleaner := CreateImportCleaner(&[]pfile.Pfile{refFile, file})
	cleaner.Clean(&file)
	if !areRowsEqual(file.Rows, expRows) {
		t.Errorf("Expected <%s> but was <%s>.\n", expRows, file.Rows)
	}
}

func TestClean_ReplaceAlias(t *testing.T) {
	rows := []string{"frax.doSmth"}
	expRows := []string{"doSmth"}
	refFile := pfile.Pfile{"", []string{}, []pfile.Goimport{}, "test", "test"}
	file := pfile.Pfile{"", rows, []pfile.Goimport{pfile.CreateGoimport("frax", "test")}, "", ""}
	cleaner := CreateImportCleaner(&[]pfile.Pfile{refFile, file})
	cleaner.Clean(&file)
	if !areRowsEqual(file.Rows, expRows) {
		t.Errorf("Expected <%s> but was <%s>.\n", expRows, file.Rows)
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