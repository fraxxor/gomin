package pfile

import (
	"de.fraxxor.gofrax/gomin/input/gofilereader"
	"testing"
)

func TestProcessGofile_Normal(t *testing.T) {
	gofile := gofilereader.Gofile{AbsolutePath: "test/testfolder/testpackage/test.go", Rows: []string{"package testpackage", "12345"}}
	processor := new(ProcessGofileImpl)
	pfile := processor.ProcessGofile(&gofile)
	if (*pfile).Package != "testpackage" {
		t.Errorf("Expected <testpackage> but was <%s>.\n", (*pfile).Package)
	}
}