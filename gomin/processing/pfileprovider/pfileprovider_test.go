package pfileprovider

import (
	"testing"
	"de.fraxxor.gofrax/gomin/input/gofilereader"
	"de.fraxxor.gofrax/gomin/processing/pfile"
)

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
	// TODO
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

type PfileprocessorSniffer struct {
	called int
}

func (sniffer *PfileprocessorSniffer) ProcessGofile(gofile *gofilereader.Gofile) (*pfile.Pfile, error) {
	sniffer.called = sniffer.called + 1
	return nil, nil
}