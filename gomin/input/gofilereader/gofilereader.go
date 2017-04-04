package gofilereader

import(
	"os"
	"bufio"
)

type GofilereaderFS struct {
	
}

func (r *GofilereaderFS) ContentOf(file string) (*Gofile, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	content := readFrom(f)
	(*content).AbsolutePath = file
	f.Close()
	return content, nil
}

func readFrom(file *os.File) *Gofile {
	buf := Gofile{Rows: make([]string, 0)}
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

