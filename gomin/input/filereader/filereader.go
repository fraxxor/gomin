package filereader

import(
	"os"
	"bufio"
)

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

func ContentOf(file string) (*Gofile, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	content := readFrom(f)
	f.Close()
	return content, nil
}

func readFrom(file *os.File) *Gofile {
	buf := Gofile{make([]string, 0)}
	reader := bufio.NewReader(file)
	var row []byte
	var err error
	for err == nil {
		row, _, err = reader.ReadLine()
		buf.Rows = append(buf.Rows, string(row))
	}
	return &buf
}

