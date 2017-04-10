package pcleaner

import (
	"testing"
	"strings"
	"de.fraxxor.gofrax/gomin/processing/pfile"
)

func TestClean_RemovePrefix(t *testing.T) {
	relativePath := "project/util/calc"
	rootPath := "D:/Test/go/src"
	file := pfile.Pfile{"", []string{}, []pfile.Goimport{}, "", rootPath + "/" + relativePath}
	cleaner := CreatePackagePathCleaner(rootPath)
	cleaner.Clean(&file)
	if !strings.EqualFold(relativePath, file.PackageAbsolutePath) {
		t.Errorf("Expected <%s> but was <%s>.\n", relativePath, file.PackageAbsolutePath)
	}
}