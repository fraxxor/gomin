package gofilecollector

import(
	"de.fraxxor.gofrax/gomin/input/gofilereader"
	"de.fraxxor.gofrax/gomin/input/godirectoryreader"
)

type Gofilecollector interface {
	CollectRecursive(enterPath string) []gofilereader.Gofile
}

type collectorEOF struct {

}

func (e collectorEOF) Error() string {
	return "Recursive EOF"
}

type GofilecollectorImpl struct{
	dirReader *godirectoryreader.Godirectoryreader
	fileReader *gofilereader.Gofilereader
}

func MakeGofilecollector(
	dirReader *godirectoryreader.Godirectoryreader,
	fileReader *gofilereader.Gofilereader) *GofilecollectorImpl {
	return &GofilecollectorImpl{dirReader, fileReader}
}

func (collector *GofilecollectorImpl) CollectRecursive(enterPath string) []gofilereader.Gofile {
	collectorClosure := collectorClosure(collector, enterPath)
	var collectedFiles []gofilereader.Gofile
	var err error
	for err = nil; err == nil; collectedFiles, err = collectorClosure() {
	}
	return collectedFiles
}

func collectorClosure(collector *GofilecollectorImpl, enterPath string) func() ([]gofilereader.Gofile, error) {
	gofiles := make([]gofilereader.Gofile, 0)
	openPaths := []string{enterPath}
	return func() ([]gofilereader.Gofile, error) {
		dir, err := (*collector.dirReader).ReadDirectory(openPaths[0])
		openPaths = openPaths[1:]
		if err != nil {
			panic(err)
		}
		for _, file := range dir.Filepaths {
			gofile, err := (*collector.fileReader).ContentOf(file)
			if (err != nil) {
				panic(err)
			}
			gofiles = append(gofiles, *gofile)
		}
		for _, filedir := range dir.Directorypaths {
			openPaths = append(openPaths, filedir)
		}
		var errorReturn error
		if len(openPaths) == 0 {
			errorReturn = new(collectorEOF)
		} else {
			errorReturn = nil
		}
		return gofiles, errorReturn
	}
}