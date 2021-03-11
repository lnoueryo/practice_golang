package main

import (
	"io"
	"log"
	"os"
)

func LoggingSettings(logFile string){
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multilogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multilogFile)
}


func main(){
	LoggingSettings("test.log")
	_, err := os.Open("abc.log")
	if err != nil{
		log.Fatalln("abc.logなんてないでほんまに")
	}
}
