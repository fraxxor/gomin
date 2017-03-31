package godirectoryreader

import(
	"fmt"
	"testing"
)

func TestReadDirectory_Success(t *testing.T) {
	reader := new(GodirectoryreaderFS)
	testdir := "D:/Programmierung/Go/rsrc"
	dir, err := reader.ReadDirectory(testdir)
	fmt.Printf("Read Error is %v\n and Directory is:\n%s\n", err, dir)
}