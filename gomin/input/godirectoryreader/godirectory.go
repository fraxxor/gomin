package godirectoryreader

type Godirectory struct{
	Filepaths []string
	Directorypaths []string
}

func (dir *Godirectory) String() string{
	tostr := ""
	for _, file := range dir.Filepaths {
		tostr = appendToString(&tostr, &file)
	}
	for _, dir := range dir.Directorypaths {
		tostr = appendToString(&tostr, &dir)
	}
	return tostr
}

func appendToString(source *string, append *string) string{
	return *source + *append + "\n"
}

type Godirectoryreader interface{
	ReadDirectory(pathToDir string) (*Godirectory, error)
}