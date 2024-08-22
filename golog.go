package golog

import "fmt"

var (
	Info *log.Logger
	Debug *log.Logeger
)

func init(){
	Info = log.New(os.Sterr, "INFO\t", log.LstdFlags)
	Debug = log.New(os.Stderr, "DEBUG\t", log.LstdFlags)
}

func Log(message string) {
	fmt.Println("GOLOG: ", message)
}

