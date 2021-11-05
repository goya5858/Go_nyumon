package main

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}

func main() {
	LoggingSettings("test.log")
	_, err := os.Open("dont_exists_file")
	if err != nil {
		log.Fatalln("exit")
	}

	log.Println("logging!")
	log.Printf("%T, %v", "test", "test")

	log.Fatalln("error")
}
