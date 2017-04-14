package pfileprovider

import (
	"testing"
	"de.fraxxor.gofrax/gomin/input/gofilereader"
	"de.fraxxor.gofrax/gomin/processing/pfile"
	"de.fraxxor.gofrax/gomin/processing/pcleaner"
)

func TestProcessFiles_WhenNotAGofileThenLeaveItOut(t *testing.T) {
	stub := PfileprocessorErrStub{}
	var processor pfile.PfileProcessor
	processor = &stub
	provider := CreateProvider(&processor)
	gofiles := []gofilereader.Gofile{gofilereader.Gofile{}, gofilereader.Gofile{}}
	provider.ProcessFiles(&gofiles)
	returnFiles := provider.GetFiles()
	if len(returnFiles) > 0 {
		t.Errorf("Expected no productive File but found %d.\n", len(returnFiles))
	}
}

func TestProcessFiles_CallToProcessorTwice(t *testing.T) {
	sniffer := PfileprocessorSniffer{}
	var processor pfile.PfileProcessor
	processor = &sniffer
	provider := CreateProvider(&processor)
	gofiles := []gofilereader.Gofile{gofilereader.Gofile{}, gofilereader.Gofile{}}
	provider.ProcessFiles(&gofiles)
	if sniffer.called != 2 {
		t.Errorf("Expected two calls to the Pfileprocessor but registered %d.\n", sniffer.called)
	}
}

func TestProcessFiles_DeliverPfilesFromProcessor(t *testing.T) {
	gofile1 := gofilereader.Gofile{}
	pfile1 := pfile.Pfile{Package: "Test"}
	stub := PfileprocessorStub{pfile1}
	var processor pfile.PfileProcessor
	processor = &stub
	provider := CreateProvider(&processor)
	provider.ProcessFiles(&[]gofilereader.Gofile{gofile1})
	returnFiles := provider.GetFiles()
	if len(returnFiles) != 1 {
		t.Errorf("Expected one pfile but were %d.\n", len(returnFiles))
		return
	}
	if (returnFiles)[0].Package != pfile1.Package {
		t.Errorf("Expected <%v> but was <%v>.\n", pfile1, (returnFiles)[0])
	}
}

func TestProcessFiles_ApplyCleaner(t *testing.T) {
	gofile1 := gofilereader.Gofile{}
	pfile1 := pfile.Pfile{Package: "Test"}
	stub := PfileprocessorStub{pfile1}
	var processor pfile.PfileProcessor
	processor = &stub
	provider := CreateProvider(&processor)
	provider.ProcessFiles(&[]gofilereader.Gofile{gofile1})
	cleanerStub := PcleanerStub{"Replaced"}
	var cleaner pcleaner.Pcleaner
	cleaner = &cleanerStub
	provider.AddCleaner(&cleaner)
	provider.CleanFiles()
	returnFiles := provider.GetFiles()
	if len(returnFiles) != 1 {
		t.Errorf("Expected one pfile but were %d.\n", len(returnFiles))
		return
	}
	if (returnFiles)[0].Package != cleanerStub.packageToApply {
		t.Errorf("Expected <%v> but was <%v>.\n", cleanerStub.packageToApply, (returnFiles)[0].Package)
	}
}

type PfileprocessorSniffer struct {
	called int
}

func (sniffer *PfileprocessorSniffer) ProcessGofile(gofile *gofilereader.Gofile) (*pfile.Pfile, error) {
	sniffer.called = sniffer.called + 1
	return &pfile.Pfile{}, nil
}

type PfileprocessorStub struct {
	returnFile pfile.Pfile
}

func (stub *PfileprocessorStub) ProcessGofile(gofile *gofilereader.Gofile) (*pfile.Pfile, error) {
	return &stub.returnFile, nil
}

type PcleanerStub struct {
	packageToApply string
}

func (stub *PcleanerStub) Clean(pfile *pfile.Pfile) {
	(*pfile).Package = stub.packageToApply
}

type PfileprocessorErrStub struct {

}

func (stub *PfileprocessorErrStub) ProcessGofile(gofile *gofilereader.Gofile) (*pfile.Pfile, error) {
	return nil, pfile.NoProductiveGoFile
}