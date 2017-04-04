package pfile

import(
	"de.fraxxor.gofrax/gomin/input/gofilereader"
)

type ProcessGofileImpl struct {
	
}

func (processor *ProcessGofileImpl) ProcessGofile(gofile *gofilereader.Gofile) *Pfile {
	pfile := Pfile{AbsolutePath: gofile.AbsolutePath, Rows: gofile.Rows}
	
	return &pfile
}