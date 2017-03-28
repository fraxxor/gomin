package filereader

type Gofile struct {
	Rows []string
}

func (gf *Gofile) String() string {
	s := ""
	for _, r := range gf.Rows {
		s = s + r + "\n"
	}
	return s
}

type Gofilereader interface {
	ContentOf(file string) (*Gofile, error)
}