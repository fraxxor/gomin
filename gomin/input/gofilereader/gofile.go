package gofilereader

type Gofile struct {
	AbsolutePath string
	Rows []string
}

func (gf *Gofile) String() string {
	s := "Gofile <" + gf.AbsolutePath + ">\n"
	for _, r := range gf.Rows {
		s = s + r + "\n"
	}
	return s
}

type Gofilereader interface {
	ContentOf(file string) (*Gofile, error)
}