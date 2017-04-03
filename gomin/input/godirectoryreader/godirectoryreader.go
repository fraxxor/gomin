package godirectoryreader

import(
	"io/ioutil"
	"strings"
)

type GodirectoryreaderFS struct{
	
}

func (r *GodirectoryreaderFS) ReadDirectory(pathToDir string) (*Godirectory, error) {
	files, err := ioutil.ReadDir(pathToDir)
	if err != nil {
		return nil, err
	}
	godir := Make()
	for _, file := range files {
		if file.IsDir() {
			godir.Directorypaths = append(godir.Directorypaths, filePathRelativeToDir(file.Name(), pathToDir))
		} else {
			godir.Filepaths = append(godir.Filepaths, filePathRelativeToDir(file.Name(), pathToDir))
		}
	}

	return &godir, nil
}

func filePathRelativeToDir(filename string, pathToDir string) string {
	if strings.HasSuffix(pathToDir, "/") || strings.HasSuffix(pathToDir, "\\") {
		return pathToDir + filename
	}
	return pathToDir + "/" + filename
}