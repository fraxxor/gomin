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
	err := processPackageInformation(gofile, &pfile)
	if err != nil {
		return nil, err
	}
	return &pfile, nil
}

func processPackageInformation(gofile *gofilereader.Gofile, pfile *Pfile) error {
	packageOrNil, rowWithPackageDecl := getPackageOrNil(gofile)
	if packageOrNil == nil {
		return new (NoPackageError)
	}
	(*pfile).Package = *packageOrNil 
	(*pfile).Rows = *deleteIndexOfSlice(&(*pfile).Rows, rowWithPackageDecl)
	packageAbsolutePathOrNil := getPackageAbsolutePathOrNil(gofile.AbsolutePath, *packageOrNil)
	if packageAbsolutePathOrNil == nil {
		return new (InvalidAbsolutePackagePath)
	}
	(*pfile).PackageAbsolutePath = *packageAbsolutePathOrNil
	return nil
}

func deleteIndexOfSlice(slice *[]string, index int) *[]string {
	sliceWithDelete := append((*slice)[:index], (*slice)[index + 1:]...)
	return &sliceWithDelete
}

func getPackageOrNil(gofile *gofilereader.Gofile) (*string, int) {
	for rowIdx, row := range (*gofile).Rows {
		if strings.HasPrefix(strings.TrimSpace(row), "package ") {
			packageName := strings.TrimPrefix(strings.TrimSpace(row), "package ")
			return &packageName, rowIdx
		}
	}
	return nil, -1
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

func processImportInformation(gofile *gofilereader.Gofile, pfile *Pfile) {

}