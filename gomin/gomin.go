package main

import (
	"fmt"
	"de.fraxxor.gofrax/gomin/input/gofilereader"
	"de.fraxxor.gofrax/gomin/input/godirectoryreader"
	"de.fraxxor.gofrax/gomin/input/gofilecollector"
	"de.fraxxor.gofrax/gomin/processing/pfile"
)

func main() {
	fmt.Println("Hello World")
	var reader gofilereader.Gofilereader
	reader = new(gofilereader.GofilereaderFS)
	content, err := reader.ContentOf("D:/Programmierung/Go/rsrc/testfile.txt")
	if err != nil {
		fmt.Printf("Error on Read: %v", err)
	} else {
		fmt.Printf("Successfully read the following:\n%s\n", content)
	}
	dirReader := new(godirectoryreader.GodirectoryreaderFS)
	godir, err := dirReader.ReadDirectory("D:/Programmierung/Go/rsrc")
	if err != nil {
		fmt.Printf("Error on DirRead: %v", err)
	} else {
		fmt.Printf("Directory = \n%s\n", godir)
	}
	var dirReaderAbstr godirectoryreader.Godirectoryreader
	dirReaderAbstr = dirReader
	dirCollector := gofilecollector.MakeGofilecollector(&dirReaderAbstr, &reader)
	gofilesRecursive := dirCollector.CollectRecursive("D:/Programmierung/Go/src/de.fraxxor.gofrax/gomin")
	var processor pfile.PfileProcessor
	processor = new(pfile.ProcessGofileImpl)
	for _, gf := range gofilesRecursive {
		fmt.Printf("Gofile: %s\n", gf.AbsolutePath)
		processedFile, err := processor.ProcessGofile(&gf)
		if (err != nil) {
			panic(err)
		}
		fmt.Printf("Pfile: %s\n", (*processedFile).PackageAbsolutePath)
	}
}
