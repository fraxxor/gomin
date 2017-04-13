package pfileprovider

import (
	"de.fraxxor.gofrax/gomin/input/gofilereader"
	"de.fraxxor.gofrax/gomin/processing/pfile"
	"de.fraxxor.gofrax/gomin/processing/pcleaner"
)

type Pfileprovider interface {
	AddCleaner(cleaner *pcleaner.Pcleaner)
	ProcessFiles(gofiles *[]gofilereader.Gofile)
	CleanFiles()
	GetFiles() []*pfile.Pfile
}

type PfileproviderImpl struct {
	processor *pfile.PfileProcessor
	cleaners []*pcleaner.Pcleaner
	pfiles []*pfile.Pfile
}

func CreateProvider(processor *pfile.PfileProcessor) *PfileproviderImpl {
	pfiles := make([]*pfile.Pfile, 0)
	return &PfileproviderImpl{processor, make([]*pcleaner.Pcleaner, 0), pfiles}
}

func (provider *PfileproviderImpl) AddCleaner(cleaner *pcleaner.Pcleaner) {
	provider.cleaners = append(provider.cleaners, cleaner)
}

func (provider *PfileproviderImpl) ProcessFiles(gofiles *[]gofilereader.Gofile) {
	for _, gofile := range *gofiles {
		pfile, err := (*provider.processor).ProcessGofile(&gofile)
		if err != nil {
			panic(err)
		}
		provider.pfiles = append(provider.pfiles, pfile)
	}
}

func (provider *PfileproviderImpl) CleanFiles() {
	for _, pfile := range provider.pfiles {
		for _, cleaner := range provider.cleaners {
			(*cleaner).Clean(pfile)
		}
	}
}

func (provider *PfileproviderImpl) GetFiles() []*pfile.Pfile {
	return provider.pfiles
}