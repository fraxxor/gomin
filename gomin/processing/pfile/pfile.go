package pfile

import(
	"de.fraxxor.gofrax/gomin/input/gofilereader"
	"strings"
)

var(
	NoProductiveGoFile = new(noProductiveGoFile)
)

type Goimport struct {
	prefix string
	importPath string
}

func CreateGoimport(prefix, importpath string) Goimport {
	if strings.HasPrefix(importpath, "\"") || strings.HasSuffix(importpath, "\"") {
		panic("Importpath must be without quotations.")
	}
	return Goimport{prefix, importpath}
}

func (g Goimport) HasPrefix() bool {
	return g.prefix != ""
}

func (g Goimport) Prefix() string {
	return g.prefix
}

func (g Goimport) ImportPath() string {
	return g.importPath
}

func (g Goimport) String() string {
	if g.HasPrefix() {
		if strings.EqualFold(getLastElementOfImportPath(g.ImportPath()), g.Prefix()) {
			return "\"" + g.ImportPath() + "\""
		}
		return g.Prefix() + " \"" + g.ImportPath() + "\""
	}
	return ". \"" + g.ImportPath() + "\""
}

type Pfile struct {
	AbsolutePath string
	Rows []string
	Imports []Goimport
	Package string
	PackageAbsolutePath string
}

type noProductiveGoFile struct {

}

func (err noProductiveGoFile) Error() string {
	return "Not a productive Go file."
}

type PfileProcessor interface {
	ProcessGofile(gofile *gofilereader.Gofile) (*Pfile, error)
}

func getLastElementOfImportPath(importpath string) string {
	elements := strings.Split(strings.Replace(importpath, "\"", "", -1), ".")
	lastElement := importpath
	for _, element := range elements {
		lastElement = element
	}
	return lastElement
}