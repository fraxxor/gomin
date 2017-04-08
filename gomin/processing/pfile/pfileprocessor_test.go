package pfile

import (
	"de.fraxxor.gofrax/gomin/input/gofilereader"
	"testing"
	"strings"
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

func TestProcessGofile_Package(t *testing.T) {
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
	expectedRows := []string{"12345"}
	if !areRowsEqual(expectedRows, (*pfile).Rows) {
		t.Errorf("Expected <%s>, but was <%s>.\n", expectedRows, (*pfile).Rows)
	}
}

func TestProcessGofile_Import(t *testing.T) {
	gofile := gofilereader.Gofile{	AbsolutePath: "test/testfolder/testpackage/test.go",
											Rows: []string{"package testpackage", "import \"frax\"", "frax.doSmth()"}}
	processor := new(ProcessGofileImpl)
	pfile, err := processor.ProcessGofile(&gofile)
	if err != nil {
		t.Errorf("Expected no error.\n")
		return
	}
	if len((*pfile).Imports) != 1 {
		t.Errorf("Expected one import, but was %d.\n", len((*pfile).Imports))
		return
	}
	if !strings.EqualFold((*pfile).Imports[0].ImportPath(), "\"frax\"") {
		t.Errorf("Expected <%s> but was <%s>.\n", "\"frax\"", ((*pfile).Imports[0]).ImportPath())
	}
	if !strings.EqualFold((*pfile).Imports[0].Prefix(), "frax") {
		t.Errorf("Expected <%s> but was <%s>.\n", "frax", ((*pfile).Imports[0]).Prefix())
	}
	expectedRows := []string{"frax.doSmth()"}
	if !areRowsEqual(expectedRows, (*pfile).Rows) {
		t.Errorf("Expected <%s>, but was <%s>.\n", expectedRows, (*pfile).Rows)
	}
}

func TestProcessGofile_ImportDot(t *testing.T) {
	gofile := gofilereader.Gofile{	AbsolutePath: "test/testfolder/testpackage/test.go",
											Rows: []string{"package testpackage", "import . \"frax\"", "doSmth()"}}
	processor := new(ProcessGofileImpl)
	pfile, err := processor.ProcessGofile(&gofile)
	if err != nil {
		t.Errorf("Expected no error.\n")
		return
	}
	if len((*pfile).Imports) != 1 {
		t.Errorf("Expected one import, but was %d.\n", len((*pfile).Imports))
		return
	}
	if !strings.EqualFold((*pfile).Imports[0].ImportPath(), "\"frax\"") {
		t.Errorf("Expected <%s> but was <%s>.\n", "\"frax\"", ((*pfile).Imports[0]).ImportPath())
	}
	if ((*pfile).Imports[0].HasPrefix()) {
		t.Errorf("Expected no prefix but was <%s>.\n", ((*pfile).Imports[0]).Prefix())
	}
	expectedRows := []string{"doSmth()"}
	if !areRowsEqual(expectedRows, (*pfile).Rows) {
		t.Errorf("Expected <%s>, but was <%s>.\n", expectedRows, (*pfile).Rows)
	}
}

func TestProcessGofile_ImportBrackets(t *testing.T) {
	gofile := gofilereader.Gofile{	AbsolutePath: "test/testfolder/testpackage/test.go",
											Rows: []string{"package testpackage", "import(", "alias \"frax\"", ")", "alias.doSmth()"}}
	processor := new(ProcessGofileImpl)
	pfile, err := processor.ProcessGofile(&gofile)
	if err != nil {
		t.Errorf("Expected no error.\n")
		return
	}
	if len((*pfile).Imports) != 1 {
		t.Errorf("Expected one import, but was %d: %s\n", len((*pfile).Imports), (*pfile).Imports)
		return
	}
	if !strings.EqualFold((*pfile).Imports[0].ImportPath(), "\"frax\"") {
		t.Errorf("Expected <%s> but was <%s>.\n", "\"frax\"", ((*pfile).Imports[0]).ImportPath())
	}
	if !strings.EqualFold((*pfile).Imports[0].Prefix(), "alias") {
		t.Errorf("Expected <%s> but was <%s>.\n", "\"alias\"", ((*pfile).Imports[0]).Prefix())
	}
	expectedRows := []string{"alias.doSmth()"}
	if !areRowsEqual(expectedRows, (*pfile).Rows) {
		t.Errorf("Expected <%s>, but was <%s>.\n", expectedRows, (*pfile).Rows)
	}
}

func TestProcessGofile_ImportEverything(t *testing.T) {
	gofile := gofilereader.Gofile{	AbsolutePath: "test/testfolder/testpackage/test.go",
											Rows: []string{
												"package testpackage",
												"import(. \"apfel\"",
												"alias \"frax\")",
												"import \"test\"",
												"alias.doSmth()",
												"test.doAnything()"}}
	processor := new(ProcessGofileImpl)
	pfile, err := processor.ProcessGofile(&gofile)
	if err != nil {
		t.Errorf("Expected no error.\n")
		return
	}
	if len((*pfile).Imports) != 3 {
		t.Errorf("Expected thre imports, but was %d: %s\n", len((*pfile).Imports), (*pfile).Imports)
		return
	}
	// Assert apfel
	if !strings.EqualFold((*pfile).Imports[0].ImportPath(), "\"apfel\"") {
		t.Errorf("Expected <%s> but was <%s>.\n", "\"apfel\"", ((*pfile).Imports[0]).ImportPath())
	}
	if ((*pfile).Imports[0].HasPrefix()) {
		t.Errorf("Expected no prefix but was <%s>.\n", ((*pfile).Imports[0]).Prefix())
	}
	// Assert frax
	if !strings.EqualFold((*pfile).Imports[1].Prefix(), "alias") {
		t.Errorf("Expected <%s> but was <%s>.\n", "alias", ((*pfile).Imports[1]).Prefix())
	}
	if !strings.EqualFold((*pfile).Imports[1].ImportPath(), "\"frax\"") {
		t.Errorf("Expected <%s> but was <%s>.\n", "\"frax\"", ((*pfile).Imports[1]).ImportPath())
	}
	// Assert test
	if !strings.EqualFold((*pfile).Imports[2].Prefix(), "test") {
		t.Errorf("Expected <%s> but was <%s>.\n", "test", ((*pfile).Imports[2]).Prefix())
	}
	if !strings.EqualFold((*pfile).Imports[2].ImportPath(), "\"test\"") {
		t.Errorf("Expected <%s> but was <%s>.\n", "\"test\"", ((*pfile).Imports[2]).ImportPath())
	}
	expectedRows := []string{"alias.doSmth()", "test.doAnything()"}
	if !areRowsEqual(expectedRows, (*pfile).Rows) {
		t.Errorf("Expected <%s>, but was <%s>.\n", expectedRows, (*pfile).Rows)
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
