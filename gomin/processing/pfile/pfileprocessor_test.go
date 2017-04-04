package pfile

import (
	"de.fraxxor.gofrax/gomin/input/gofilereader"
	"testing"
)

func TestProcessGofile_NoPackage(t *testing.T) {
	gofile := gofilereader.Gofile{	AbsolutePath: "",
											Rows: []string{"12345"}}
	processor := new(ProcessGofileImpl)
	_,err := processor.ProcessGofile(&gofile)
	if err == nil {
		t.Errorf("Expected error.\n")
	}
}

func TestProcessGofile_InvalidPackageAbsolutePath(t *testing.T) {
	gofile := gofilereader.Gofile{	AbsolutePath: "apple/test/test.go",
											Rows: []string{"package apple", "12345"}}
	processor := new(ProcessGofileImpl)
	_,err := processor.ProcessGofile(&gofile)
	if err == nil {
		t.Errorf("Expected error.\n")
	}
}

func TestProcessGofile_Normal(t *testing.T) {
	gofile := gofilereader.Gofile{	AbsolutePath: "test/testfolder/testpackage/test.go",
											Rows: []string{"package testpackage", "12345"}}
	processor := new(ProcessGofileImpl)
	pfile, err := processor.ProcessGofile(&gofile)
	if err != nil {
		t.Errorf("Expected no error.\n")
		return
	}
	if (*pfile).Package != "testpackage" {
		t.Errorf("Expected <testpackage> but was <%s>.\n", (*pfile).Package)
	}
	if (*pfile).PackageAbsolutePath != "test/testfolder/testpackage" {
		t.Errorf("Expected <test/testfolder/testpackage> but was <%s>.\n", (*pfile).PackageAbsolutePath)	
	}
}