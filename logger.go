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

//Appends err to file
func ErrLogger(e error, Fname string) {
	file, err := os.OpenFile(Fname, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	check(err)
	defer file.Close()
	log.SetOutput(file)
	log.Println(e)

}
