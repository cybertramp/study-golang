package main

import(
	_ "fmt"
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main(){
	programName := filepath.Base(os.Args[0])
	sysLog, err := syslog.New(syslog.LOG_INFO | syslog.LOG_LOCAL7, programName)

	if err != nil{
		log.Fatal(err)		// err message to Fatal
	} else{
		log.SetOutput(sysLog)
	}
	log.Println("LOG_INFO + LOG_LOCAL7: Logging in Go!")
}