package pmerger

import (
	"testing"
	"strings"
	"de.fraxxor.gofrax/gomin/processing/pfile"
)

func TestMerge_Empty(t *testing.T) {
	merger := CreateMerger()
	mergefile := merger.Merge(&[]pfile.Pfile{})
	if len(mergefile.Rows) > 0 {
		t.Errorf("Expected empty rows but was <%s>.\n", mergefile.Rows)
	}
	if mergefile.Package != "main" {
		t.Errorf("Expected <\"main\"> but was <%s>.\n", mergefile.Package)
	}
}

func TestMerge_Single(t *testing.T) {
	merger := CreateMerger()
	singleFile := pfile.Pfile{Rows: []string{"Test"}, Imports: []pfile.Goimport{pfile.CreateGoimport("imp", "imp")}}
	mergefile := merger.Merge(&[]pfile.Pfile{singleFile})
	if len(mergefile.Rows) != 1 {
		t.Errorf("Expected one row but was <#%d: %s>.\n", len(mergefile.Rows), mergefile.Rows)
	}
	if len(mergefile.Imports) != 1 {
		t.Errorf("Expected one import but was <%s>.\n", mergefile.Imports)	
	}
}

func TestString(t *testing.T) {
	mergefile := Mergefile{
		Package: "main",
		Imports: []pfile.Goimport{
			pfile.CreateGoimport("test", "test"),
			pfile.CreateGoimport("alias", "frax"),
			pfile.CreateGoimport("", "global"),
		},
		Rows: []string{
			"First row",
			"Second row",
		},
	}
	expected := "package main"
	expected = expected + "\n" + "import \"test\""
	expected = expected + "\n" + "import alias \"frax\""
	expected = expected + "\n" + "import . \"global\""
	expected = expected + "\n" + "First row"
	expected = expected + "\n" + "Second row"
	if !strings.EqualFold(expected, mergefile.String()) {
		t.Errorf("Expected <%s>\nbut was <%s>.\n", expected, mergefile.String())
	}
}