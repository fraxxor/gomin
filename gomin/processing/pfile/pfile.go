package pfile

import(
	"de.fraxxor.gofrax/gomin/input/gofilereader"
)

type Goimport struct {
	prefix string
	importPath string
}

func CreateGoimport(prefix, importpath string) Goimport {
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
		return g.Prefix() + " " + g.ImportPath()
	}
	return g.ImportPath()
}

type Pfile struct {
	AbsolutePath string
	Rows []string
	Imports []Goimport
	Package string
	PackageAbsolutePath string
}

type PfileProcessor interface {
	ProcessGofile(gofile *gofilereader.Gofile) (*Pfile, error)
}