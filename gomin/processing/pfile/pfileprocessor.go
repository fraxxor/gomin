package pfile

import(
	"de.fraxxor.gofrax/gomin/input/gofilereader"
	"strings"
)

type ProcessGofileImpl struct {
	
}

type NoPackageError struct {

}

func (e NoPackageError) Error() string {
	return "No package declaration found."
}

type InvalidAbsolutePackagePath struct {

}

func (e InvalidAbsolutePackagePath) Error() string {
	return "Invalid absolute package path."
}

func (processor *ProcessGofileImpl) ProcessGofile(gofile *gofilereader.Gofile) (*Pfile, error) {
	pfile := Pfile{AbsolutePath: gofile.AbsolutePath, Rows: gofile.Rows}
	packageOrNil := getPackageOrNil(gofile)
	if packageOrNil == nil {
		return nil, new (NoPackageError)
	}
	pfile.Package = *packageOrNil
	packageAbsolutePathOrNil := getPackageAbsolutePathOrNil(gofile.AbsolutePath, *packageOrNil)
	if packageAbsolutePathOrNil == nil {
		return nil, new (InvalidAbsolutePackagePath)
	}
	pfile.PackageAbsolutePath = *packageAbsolutePathOrNil
	return &pfile, nil
}

func getPackageOrNil(gofile *gofilereader.Gofile) *string {
	for _, row := range (*gofile).Rows {
		if strings.HasPrefix(strings.TrimSpace(row), "package ") {
			packageName := strings.TrimPrefix(strings.TrimSpace(row), "package ")
			return &packageName
		}
	}
	return nil
}

func getPackageAbsolutePathOrNil(absolutePath string, packageName string) *string {
	canonicalPath := strings.Replace(absolutePath, "\\", "/", -1)
	pathElements := strings.Split(canonicalPath, "/")
	if !strings.HasSuffix(pathElements[len(pathElements) - 1], ".go") {
		return nil
	}
	if pathElements[len(pathElements) - 2] != packageName {
		return nil
	}
	packageAbsolutePath := strings.TrimSuffix(canonicalPath, "/" + pathElements[len(pathElements) - 1])
	return &packageAbsolutePath
}