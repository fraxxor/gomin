package main

import (
	"fmt"
	"os"
    "path/filepath"
	"de.fraxxor.gofrax/gomin/input/gofilereader"
	"de.fraxxor.gofrax/gomin/input/godirectoryreader"
	"de.fraxxor.gofrax/gomin/input/gofilecollector"
	"de.fraxxor.gofrax/gomin/processing/pfile"
	"de.fraxxor.gofrax/gomin/processing/pfileprovider"
	"de.fraxxor.gofrax/gomin/processing/pmerger"
	"de.fraxxor.gofrax/gomin/processing/pcleaner"
)

func main() {
	pathOfExec := getRootPath()
	pathOfExec = "D:/Programmierung/Go/src/de.fraxxor.gofrax/gomin"
	pathToRoot := "D:/Programmierung/Go/src"
	
	var reader gofilereader.Gofilereader
	reader = new(gofilereader.GofilereaderFS)

	var dirReader godirectoryreader.Godirectoryreader
	dirReader = new(godirectoryreader.GodirectoryreaderFS)

	var dirCollector gofilecollector.Gofilecollector
	dirCollector = gofilecollector.MakeGofilecollector(&dirReader, &reader)

	gofilesRecursive := dirCollector.CollectRecursive(pathOfExec)

	var processor pfile.PfileProcessor
	processor = new(pfile.ProcessGofileImpl)

	var provider pfileprovider.Pfileprovider
	provider = pfileprovider.CreateProvider(&processor)

	provider.ProcessFiles(&gofilesRecursive)
	processedPfiles := provider.GetFiles()

	var packagecleaner pcleaner.Pcleaner
	packagecleaner = pcleaner.CreatePackagePathCleaner(pathToRoot)

	var importcleaner pcleaner.Pcleaner
	importcleaner = pcleaner.CreateImportCleaner(processedPfiles)

	provider.AddCleaner(&packagecleaner)
	provider.AddCleaner(&importcleaner)

	provider.CleanFiles()
	cleanedFiles := provider.GetFiles()

	var merger pmerger.Pmerger
	merger = pmerger.CreateMerger()

	mergefile := merger.Merge(cleanedFiles)
	fmt.Printf("Mergefile:\n%s\n", mergefile)
}

func getRootPath() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	return dir
}
