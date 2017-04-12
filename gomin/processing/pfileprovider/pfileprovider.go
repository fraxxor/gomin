package pfileprovider

import (
	"de.fraxxor.gofrax/gomin/input/gofilereader"
	"de.fraxxor.gofrax/gomin/processing/pfile"
	"de.fraxxor.gofrax/gomin/processing/pcleaner"
)

type Pfileprovider interface {
	AddCleaner(cleaner *pcleaner.Pcleaner)
	ProcessFiles(gofiles *[]gofilereader.Gofile) *[]pfile.Pfile
}

type PfileproviderImpl struct {
	processor *pfile.PfileProcessor
	cleaners []*pcleaner.Pcleaner
}

func CreateProvider(processor *pfile.PfileProcessor) *PfileproviderImpl {
	return &PfileproviderImpl{processor, make([]*pcleaner.Pcleaner, 0)}
}

func (provider *PfileproviderImpl) AddCleaner(cleaner *pcleaner.Pcleaner) {

}

func (provider *PfileproviderImpl) ProcessFiles(gofiles *[]gofilereader.Gofile) *[]pfile.Pfile {
	pfiles := make([]pfile.Pfile, 0)
	for _, gofile := range *gofiles {
		pfile, err := (*provider.processor).ProcessGofile(&gofile)
		if err != nil {
			panic(err)
		}
		pfiles = append(pfiles, *pfile)
	}
	return &pfiles
}