package godirectoryreader

import(
	"fmt"
	"testing"
	"strings"
)

func TestReadDirectory_Success(t *testing.T) {
	reader := new(GodirectoryreaderFS)
	testdir := "D:/Programmierung/Go/rsrc"
	dir, err := reader.ReadDirectory(testdir)
	fmt.Printf("Read Error is %v\n and Directory is:\n%s\n", err, dir)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	for _, file := range dir.Filepaths {
		if !strings.HasPrefix(file, testdir) {
			t.Errorf("Expected path to start with <%s>, but was <%s>", testdir, file)
		}
	}
	for _, dirpath := range dir.Directorypaths {
		if !strings.HasPrefix(dirpath, testdir) {
			t.Errorf("Expected path to start with <%s>, but was <%s>", testdir, dirpath)
		}
	}
}