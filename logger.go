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

//Appends err to filei
func ErrLogger(e error, Fname string) {

	file, err := os.OpenFile(Fname, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	check(err)
	defer file.Close()
	logger := log.New(file, "prefix", log.LstdFlags|log.Lshortfile)
	logger.SetOutput(file)
	logger.Println(e)

}
