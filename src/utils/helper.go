package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Info Log Function exported
var Info *log.Logger

// Error Log Function exported
var Error *log.Logger

// Debug Log function exported
var Debug *log.Logger

func init() {
	absPath, err := filepath.Abs("../log")
	if err != nil {
		fmt.Println("Errore nel leggere il path:", err)
	}

	generalLog, err := os.OpenFile(absPath+"/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Errore nell'apertura del file:", err)
		os.Exit(1)
	}
	Info = log.New(generalLog, "[INFO]:\t", log.Ldate|log.Ltime|log.Lshortfile)
	Debug = log.New(generalLog, "[DEBUG]:\t", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(generalLog, "[ERROR]:\t", log.Ldate|log.Ltime|log.Lshortfile)
}
