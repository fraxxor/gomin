package godirectoryreader

import(
	"io/ioutil"
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
			godir.Directorypaths = append(godir.Directorypaths, file.Name())
		} else {
			godir.Filepaths = append(godir.Filepaths, file.Name())
		}
	}

	return &godir, nil
}