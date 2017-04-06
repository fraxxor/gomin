package pcleaner

import (
	"de.fraxxor.gofrax/gomin/processing/pfile"
)

type Pcleaner interface {
	Clean(pfile *pfile.Pfile)
}