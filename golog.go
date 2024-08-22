package golog

import (
	"fmt"
)

func Info(message string) {
	fmt.Println("INFO: ", message)
}

func Debug(message string) {
	fmt.Println("DEBUG: ", message)
}

