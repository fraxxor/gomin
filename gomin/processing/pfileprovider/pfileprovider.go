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
	for _, gofile := range *gofiles {
		_, err := (*provider.processor).ProcessGofile(&gofile)
		if (err != nil) {
			panic(err)
		}
	}
	return nil
}