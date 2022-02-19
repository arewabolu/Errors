package logger

import (
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		log.Panic("could not open file")
	}
}

func ErrLogger(t string) {
	buf := []byte(t)
	file, err := os.OpenFile("Errorlog", os.O_CREATE|os.O_APPEND, 0644)
	check(err)
	defer file.Close()

	_, err2 := file.WriteAt(buf, int64(len(buf)))
	check(err2)

}
