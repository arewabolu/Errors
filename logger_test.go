package logger

import (
	"errors"
	"os"
	"testing"

	_ "github.com/arewabolu/Errors"
)

func TestErrlogger(t *testing.T) {
	str := "No errors when logging"
	err := errors.New(str)
	ErrLogger(err, "file.txt")
	file, _ := os.OpenFile("file.txt", os.O_CREATE|os.O_RDONLY, 0644)

	buf := make([]byte, int(len(str)))

	defer file.Close()

	size, _ := file.Read(buf)
	if size != int(len(str)) {
		t.Error("Something wrong in reading the file")
	}
}
