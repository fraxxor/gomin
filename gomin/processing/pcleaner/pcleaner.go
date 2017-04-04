package pcleaner

import (
	"de.fraxxor.gofrax/gomin/processing/pfile"
)

type Pcleaner interface {
	CleanFromImports(pfile *pfile.Pfile) *pfile.Pfile
}

type PcleanerImpl struct {

}

func (cleaner *PcleanerImpl) CleanFromImports(pfile *pfile.Pfile) *pfile.Pfile {
	return pfile
}