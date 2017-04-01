package gofilecollector

import(
	"de.fraxxor.gofrax/gomin/input/gofilereader"
	"de.fraxxor.gofrax/gomin/input/godirectoryreader"
)

type Gofilecollector struct{
	dirReader *godirectoryreader.Godirectoryreader
	fileReader *gofilereader.Gofilereader
}

func MakeGofilecollector(
	dirReader *godirectoryreader.Godirectoryreader,
	fileReader *gofilereader.Gofilereader) *Gofilecollector {
	return &Gofilecollector{dirReader, fileReader}
}

func (collector *Gofilecollector) CollectRecursive(enterPath string) []gofilereader.Gofile {
	gofiles := make([]gofilereader.Gofile, 0)
	dir, err := (*collector.dirReader).ReadDirectory(enterPath)
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
	return gofiles
}