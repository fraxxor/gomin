package pcleaner

import(
	"de.fraxxor.gofrax/gomin/processing/pfile"
	"strings"
)

type ImportCleaner struct {
	replacableImports *[]string
}

func CreateImportCleaner(pfiles *[]pfile.Pfile) *ImportCleaner {
	allreplacableImports := make([]string, 0)
	for _, pfile := range *pfiles {
		allreplacableImports = append(allreplacableImports, pfile.PackageAbsolutePath)
	}
	return &ImportCleaner{&allreplacableImports}
}

func (cleaner *ImportCleaner) Clean(fileToClean *pfile.Pfile) {
	importsToClean := make([]pfile.Goimport, 0)
	for _, goimport := range fileToClean.Imports {
		for _, replacable := range *cleaner.replacableImports {
			if strings.EqualFold(goimport.ImportPath(), replacable) {
				importsToClean = append(importsToClean, goimport)
				break
			}
		}
	}
	for i, row := range (*fileToClean).Rows {
		cleanedRow := row
		for _, importToClean := range importsToClean {
			cleanedRow = strings.Replace(cleanedRow, importToClean.Prefix() + ".", "", -1)
		}
		(*fileToClean).Rows[i] = cleanedRow
	}
	remainingImports := make([]pfile.Goimport, 0)
	for _, remainingImport := range (*fileToClean).Imports {
		keepImport := true
		for _, importToClean := range importsToClean {
			if remainingImport == importToClean {
				keepImport = false
				break
			}
		}
		if keepImport {
			remainingImports = append(remainingImports, remainingImport)
		}
	}
	(*fileToClean).Imports = remainingImports
}