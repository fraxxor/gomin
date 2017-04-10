package pcleaner

import (
	"de.fraxxor.gofrax/gomin/processing/pfile"
	"strings"
)

type PackagePathCleaner struct {
	pathToRoot string
}

func CreatePackagePathCleaner(pathToRoot string) *PackagePathCleaner {
	return &PackagePathCleaner{strings.Replace(pathToRoot, "\\", "/", -1)}
}

func (cleaner *PackagePathCleaner) Clean(fileToClean *pfile.Pfile) {
	if strings.HasPrefix((*fileToClean).PackageAbsolutePath, cleaner.pathToRoot + "/") {
		(*fileToClean).PackageAbsolutePath = strings.TrimPrefix((*fileToClean).PackageAbsolutePath, cleaner.pathToRoot + "/")
	}
}