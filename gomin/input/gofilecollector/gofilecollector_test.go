package gofilecollector

import(
	"testing"
	"de.fraxxor.gofrax/gomin/input/gofilereader"
	"de.fraxxor.gofrax/gomin/input/godirectoryreader"
)

func TestCollectRecursive_OnlyFiles(t *testing.T) {
	var dirReaderStub godirectoryreader.Godirectoryreader
	dirReaderStub = MakeDirReader(makeDirWith([]string{"Testfile.txt", "Andere Datei.csv"}, []string{}))
	var fileReaderStub gofilereader.Gofilereader
	fileReaderStub = &FileReaderStub{gofilereader.Gofile{"", []string{"Testzeile"}}}
	collector := MakeGofilecollector(&dirReaderStub, &fileReaderStub)
	gofiles := collector.CollectRecursive("")
	if len(gofiles) != 2 {
		t.Errorf("Expected Gofiles to be 2, but were %d", len(gofiles))
	}
}

func TestCollectRecursive_OneSubDir(t *testing.T) {
	var dirReaderStub godirectoryreader.Godirectoryreader
	dirReaderStub = MakeDirReader(makeDirWith([]string{"Testfile.txt", "Andere Datei.csv"}, []string{"Subdir"}))
	var fileReaderStub gofilereader.Gofilereader
	fileReaderStub = &FileReaderStub{gofilereader.Gofile{"", []string{"Testzeile"}}}
	collector := MakeGofilecollector(&dirReaderStub, &fileReaderStub)
	gofiles := collector.CollectRecursive("")
	if len(gofiles) != 4 {
		t.Errorf("Expected Gofiles to be 4, but were %d", len(gofiles))
	}
}

func makeDirWith(files []string, dirs []string) godirectoryreader.Godirectory {
	dir := godirectoryreader.Make()
	dir.Filepaths = files
	dir.Directorypaths = dirs
	return dir
}

type DirReaderStub struct {
	amountOfDirReturns int
	amountOfMaxDirReturns int
	dir godirectoryreader.Godirectory
}

func MakeDirReader(dir godirectoryreader.Godirectory) *DirReaderStub {
	return &DirReaderStub{0, 1, dir}
}

func (s *DirReaderStub) ReadDirectory(pathToDir string) (*godirectoryreader.Godirectory, error) {
	if s.amountOfDirReturns < s.amountOfMaxDirReturns {
		s.amountOfDirReturns = s.amountOfDirReturns + 1
		return &s.dir, nil
	} else {
		dirWithoutSubdirs := makeDirWith(s.dir.Filepaths, []string{})
		return &dirWithoutSubdirs, nil
	}
}

type FileReaderStub struct {
	file gofilereader.Gofile
}

func (s *FileReaderStub) ContentOf(file string) (*gofilereader.Gofile, error) {
	return &s.file, nil
}