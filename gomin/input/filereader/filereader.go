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
	scanner := bufio.NewScanner(file)
	endOfScan := false
	for !endOfScan {
		endOfScan = !scanner.Scan()
		if scanner.Err() == nil {
			row := scanner.Text()
			buf.Rows = append(buf.Rows, row)	
		}
	}
	return &buf
}

