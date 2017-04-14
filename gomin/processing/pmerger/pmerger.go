package pmerger

import (
	"de.fraxxor.gofrax/gomin/processing/pfile"
)

type Mergefile struct {
	Rows []string
	Imports []pfile.Goimport
	Package string
}

func (file Mergefile) String() string {
	val := "package " + file.Package
	for _, goimport := range file.Imports {
		val = val + "\n" + "import " + goimport.String()
	}
	for _, row := range file.Rows {
		val = val + "\n" + row
	}
	return val
}

type Pmerger interface {
	Merge (pfiles *[]*pfile.Pfile) *Mergefile
}

type PmergerImpl struct {

}

func CreateMerger() *PmergerImpl {
	return &PmergerImpl{}
}

func (merger *PmergerImpl) Merge(pfiles *[]*pfile.Pfile) *Mergefile {
	rowsize := 0
	importsize := 0
	for _, onefile := range *pfiles {
		rowsize = rowsize + len((*onefile).Rows)
		importsize = importsize + len((*onefile).Imports)
	}
	rows := make([]string, 0, rowsize)
	imports := make([]pfile.Goimport, 0, importsize)
	for _, onefile := range *pfiles {
		for _, row := range (*onefile).Rows {
			rows = append(rows, row)
		}
		for _, goimport := range (*onefile).Imports {
			imports = append(imports, goimport)
		}
	}
	return &Mergefile{Package: "main", Rows: rows, Imports: imports}
}