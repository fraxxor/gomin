package pcleaner

import (
	"testing"
	"de.fraxxor.gofrax/gomin/processing/pfile"
)

func TestClean_Identical(t *testing.T) {
	rows := []string{"first row", "second row"}
	file := pfile.Pfile{"", rows, []pfile.Goimport{}, "", ""}
	cleaner := CreateImportCleaner(&[]pfile.Pfile{file})
	cleaner.Clean(&file)
	if !areRowsEqual(file.Rows, rows) {
		t.Errorf("Expected <%s> but was <%s>.\n", rows, file.Rows)
	}
}

func TestClean_Replaced(t *testing.T) {
	rows := []string{"import \"test\"", "test.doSmth"}
	expRows := []string{"doSmth"}
	file := pfile.Pfile{"", rows, []pfile.Goimport{pfile.CreateGoimport("test", "test")}, "", ""}
	cleaner := CreateImportCleaner(&[]pfile.Pfile{file})
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