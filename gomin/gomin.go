package main

import (
	"os"
	"flag"
    "path/filepath"
	"de.fraxxor.gofrax/gomin/input/gofilereader"
	"de.fraxxor.gofrax/gomin/input/godirectoryreader"
	"de.fraxxor.gofrax/gomin/input/gofilecollector"
	"de.fraxxor.gofrax/gomin/processing/pfile"
	"de.fraxxor.gofrax/gomin/processing/pfileprovider"
	"de.fraxxor.gofrax/gomin/processing/pmerger"
	"de.fraxxor.gofrax/gomin/processing/pcleaner"
	"github.com/atotto/clipboard"
)

func main() {
	pathToRead := flag.String("d", getExecPath(), "Absolute Directory to read files from")
	srcPath := flag.String("s", getSrcPath(), "Source Directory")
	flag.Parse()

	mergefile := produceMergefile(*pathToRead, *srcPath)

	clipboard.WriteAll(mergefile.String())
}

func getExecPath() string {
	dir, err := os.Getwd()
	if err != nil {
		dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			return ""
		}
	}
	return dir
}

func getSrcPath() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		return ""
	}
	return filepath.Join(gopath, "src")
}

func produceMergefile(pathToRead, rootPath string) *pmerger.Mergefile{	
	var reader gofilereader.Gofilereader
	reader = new(gofilereader.GofilereaderFS)

	var dirReader godirectoryreader.Godirectoryreader
	dirReader = new(godirectoryreader.GodirectoryreaderFS)

	var dirCollector gofilecollector.Gofilecollector
	dirCollector = gofilecollector.MakeGofilecollector(&dirReader, &reader)

	gofilesRecursive := dirCollector.CollectRecursive(pathToRead)

	var processor pfile.PfileProcessor
	processor = new(pfile.ProcessGofileImpl)

	var provider pfileprovider.Pfileprovider
	provider = pfileprovider.CreateProvider(&processor)

	provider.ProcessFiles(&gofilesRecursive)
	processedPfiles := provider.GetFiles()

	var packagecleaner pcleaner.Pcleaner
	packagecleaner = pcleaner.CreatePackagePathCleaner(rootPath)

	var importcleaner pcleaner.Pcleaner
	importcleaner = pcleaner.CreateImportCleaner(processedPfiles)

	provider.AddCleaner(&packagecleaner)
	provider.AddCleaner(&importcleaner)

	provider.CleanFiles()
	cleanedFiles := provider.GetFiles()

	var merger pmerger.Pmerger
	merger = pmerger.CreateMerger()

	mergefile := merger.Merge(cleanedFiles)

	return mergefile
}