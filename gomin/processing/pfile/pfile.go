package pfile

import(
	"de.fraxxor.gofrax/gomin/input/gofilereader"
)

type Pfile struct {
	AbsolutePath string
	Rows []string
	Package string
	PackageAbsolutePath string
}

type PfileProcessor interface {
	ProcessGofile(gofile *gofilereader.Gofile) *Pfile
}