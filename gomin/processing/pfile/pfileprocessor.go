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
	processImportInformation(&pfile)
	return &pfile, nil
}

func deleteIndexOfSlice(slice *[]string, index int) *[]string {
	sliceWithDelete := append((*slice)[:index], (*slice)[index + 1:]...)
	return &sliceWithDelete
}

/**
 * Determines the package and package path.
 * Removes every package declaration row.
 * Returns an error iff package and package path do not fit
 */
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

/**
 * Determines imports.
 * Removes every import row.
 */
func processImportInformation(pfile *Pfile) {
	goimports := (*pfile).Imports
	rowsToDelete := make([]int, 0)
	importcontext := false
	for rowIdxWithImport, row := range (*pfile).Rows {
		trimmed := strings.TrimSpace(row)
		if !importcontext && strings.HasPrefix(trimmed, "import") {
			statement := strings.TrimSpace(strings.TrimPrefix(trimmed, "import"))
			if strings.HasPrefix(statement, "(") {
				importcontext = true
				statement = strings.TrimSpace(strings.TrimPrefix(statement, "("))
			}
			if len(statement) > 0 {
				prefix, importpath := getContentsFromStatement(statement)
				goimport := Goimport{prefix, importpath}
				goimports = append(goimports, goimport)
			}
			rowsToDelete = append(rowsToDelete, rowIdxWithImport)
		} else if importcontext {
			if strings.HasSuffix(trimmed, ")") {
				importcontext = false
				trimmed = strings.TrimSpace(strings.TrimSuffix(trimmed, ")"))
			}
			if len(trimmed) > 0 {
				prefix, importpath := getContentsFromStatement(trimmed)
				goimport := Goimport{prefix, importpath}
				goimports = append(goimports, goimport)
			}
			rowsToDelete = append(rowsToDelete, rowIdxWithImport)
		}
	}
	for i := len(rowsToDelete) - 1; i >= 0; i = i - 1 {
		(*pfile).Rows = *deleteIndexOfSlice(&(*pfile).Rows, rowsToDelete[i])
	}
	(*pfile).Imports = goimports
}

func getContentsFromStatement(statement string) (string, string) {
	if strings.HasPrefix(statement, "\"") {
		return getLastElementOfImportPath(statement), statement
	} else if strings.HasPrefix(statement, ". ") {
		return "", strings.TrimSpace(strings.TrimPrefix(statement, ". "))
	}
	words := strings.Split(statement, " ")
	return words[0], strings.TrimSpace(strings.TrimPrefix(statement, words[0]))
}
