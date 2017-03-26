package filereader

import (
	"testing"
	"fmt"
)

func TestContentOf_Success(t *testing.T) {
	testfile := "D:/Programmierung/Go/rsrc/testfile.txt"
	content, err := ContentOf(testfile)
	if err != nil {
		t.Errorf("Expected no error, but was %v", err)
	}
	fmt.Printf("%v", content)
}

func TestContentOf_HalloDatei(t *testing.T) {
	testfile := "D:/Programmierung/Go/rsrc/testfile.txt"
	content, err := ContentOf(testfile)
	if err != nil {
		t.Errorf("Expected no error, but was %v", err)
	}
	if content.Rows[0] != "Hallo Datei." {
		t.Errorf("Expected First input to be <Hallo Datei>, but was <%s>", content.Rows[0])
	}
}