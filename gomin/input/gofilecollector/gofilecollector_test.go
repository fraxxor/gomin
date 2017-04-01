package gofilecollector

import(
	"testing"
	"de.fraxxor.gofrax/gomin/input/gofilereader"
	"de.fraxxor.gofrax/gomin/input/godirectoryreader"
)

func TestCollectRecursive_OnlyFiles(t *testing.T) {
	var dirReaderStub godirectoryreader.Godirectoryreader
	dirReaderStub = &DirReaderStub{makeDirWith([]string{"Testfile.txt", "Andere Datei.csv"}, []string{})}
	var fileReaderStub gofilereader.Gofilereader
	fileReaderStub = &FileReaderStub{gofilereader.Gofile{[]string{"Testzeile"}}}
	collector := MakeGofilecollector(&dirReaderStub, &fileReaderStub)
	gofiles := collector.CollectRecursive("")
	if len(gofiles) != 2 {
		t.Errorf("Expected Gofiles to be 2, but were %d", len(gofiles))
	}
}

func TestCollectRecursive_OneSubDir(t *testing.T) {
	var dirReaderStub godirectoryreader.Godirectoryreader
	dirReaderStub = &DirReaderStub{makeDirWith([]string{"Testfile.txt", "Andere Datei.csv"}, []string{"Subdir"})}
	var fileReaderStub gofilereader.Gofilereader
	fileReaderStub = &FileReaderStub{gofilereader.Gofile{[]string{"Testzeile"}}}
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
	dir godirectoryreader.Godirectory
}

func (s *DirReaderStub) ReadDirectory(pathToDir string) (*godirectoryreader.Godirectory, error) {
	return &s.dir, nil
}

type FileReaderStub struct {
	file gofilereader.Gofile
}

func (s *FileReaderStub) ContentOf(file string) (*gofilereader.Gofile, error) {
	return &s.file, nil
}