package pfile

import(
	"de.fraxxor.gofrax/gomin/input/gofilereader"
	"strings"
	"log"
)

type ProcessGofileImpl struct {
	
}

type NoPackageError struct {

}

func (e NoPackageError) Error() string {
	return "No package declaration found."
}

func (processor *ProcessGofileImpl) ProcessGofile(gofile *gofilereader.Gofile) (*Pfile, error) {
	if isNotAProductiveGoFile(gofile.AbsolutePath) {
		return nil, NoProductiveGoFile
	}
	pfile := Pfile{AbsolutePath: gofile.AbsolutePath, Rows: gofile.Rows}
	err := processPackageInformation(gofile, &pfile)
	if err != nil {
		return nil, err
	}
	processImportInformation(&pfile)
	return &pfile, nil
}

func isNotAProductiveGoFile(path string) bool {
	if strings.HasSuffix(path, "_test.go") {
		return true
	}
	if strings.HasSuffix(path, ".go") {
		return false
	}
	return true
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
	packageAbsolutePath := getPackageAbsolutePath(gofile.AbsolutePath, *packageOrNil)
	(*pfile).PackageAbsolutePath = *packageAbsolutePath
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

func getPackageAbsolutePath(absolutePath string, packageName string) *string {
	canonicalPath := strings.Replace(absolutePath, "\\", "/", -1)
	pathElements := strings.Split(canonicalPath, "/")
	packageAbsolutePath := strings.TrimSuffix(canonicalPath, "/" + pathElements[len(pathElements) - 1])
	if pathElements[len(pathElements) - 2] != packageName {
		log.SetPrefix("WARN ")
		log.Printf("Package <%s> differs from Filepath <%s>.\n", packageName, packageAbsolutePath)
	}
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
				goimport := CreateGoimport(prefix, cleanseImportpathFromQuotes(importpath))
				goimports = append(goimports, goimport)
			}
			rowsToDelete = append(rowsToDelete, rowIdxWithImport)
		} else if !importcontext{
			if len(trimmed) > 0 {
				// Productive row terminates imports
				break	
			}
		} else if importcontext {
			if strings.HasSuffix(trimmed, ")") {
				importcontext = false
				trimmed = strings.TrimSpace(strings.TrimSuffix(trimmed, ")"))
			}
			if len(trimmed) > 0 {
				prefix, importpath := getContentsFromStatement(trimmed)
				goimport := CreateGoimport(prefix, cleanseImportpathFromQuotes(importpath))
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

func cleanseImportpathFromQuotes(importpath string) string {
	if strings.HasPrefix(importpath, "\"") && strings.HasSuffix(importpath, "\"") {
		return strings.TrimSuffix(strings.TrimPrefix(importpath, "\""), "\"")
	}
	return importpath
}