package pcleaner

import (
	"de.fraxxor.gofrax/gomin/processing/pfile"
	"strings"
)

type Pcleaner interface {
	CleanFromImports(pfile *pfile.Pfile)
}

type PcleanerImpl struct {

}

func (cleaner *PcleanerImpl) CleanFromImports(pfile *pfile.Pfile) {
	cleanedRows := make([]string, 0)
	for _, row := range (*pfile).Rows {
		if !strings.HasPrefix(strings.TrimSpace(row), "import ") {
			cleanedRows = append(cleanedRows, row)
		}
	}
	(*pfile).Rows = cleanedRows
}