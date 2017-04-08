package pcleaner

import(
	"de.fraxxor.gofrax/gomin/processing/pfile"
	"strings"
)

type ImportCleaner struct {
	goimports *[]pfile.Goimport
}

func CreateImportCleaner(pfiles *[]pfile.Pfile) *ImportCleaner {
	allGoimports := make([]pfile.Goimport, 0)
	for _, pfile := range *pfiles {
		for _, goimport := range pfile.Imports {
			allGoimports = append(allGoimports, goimport)
		}
	}
	return &ImportCleaner{&allGoimports}
}

func (cleaner *ImportCleaner) Clean(pfile *pfile.Pfile) {
	cleanedRows := make([]string, 0)
	for _, row := range (*pfile).Rows {
		if !strings.HasPrefix(strings.TrimSpace(row), "import ") {
			cleanedRows = append(cleanedRows, row)
		}
	}
	(*pfile).Rows = cleanedRows
}