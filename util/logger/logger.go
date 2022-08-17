package logger

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	InfoLogger    *log.Logger
	DebugLogger   *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
)

func init() {

	//Comprueba la existencia de la ruta
	//En caso de error comprueba la coincidencia del error "file does not exist"
	//Si es así, crea la ruta con permisos generales
	if _, err := os.Stat("logs"); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir("logs", os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	filename, _ := createFilename()

	file, err := os.OpenFile("logs/"+filename+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.LstdFlags|log.Lshortfile)
	DebugLogger = log.New(file, "DEBUG: ", log.LstdFlags|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.LstdFlags|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.LstdFlags|log.Lshortfile)
}

func createTwoDigits(date int) (string, error) {
	if date < 10 {
		return "0" + strconv.Itoa(date), nil
	}
	return strconv.Itoa(date), nil
}

func createFilename() (string, error) {
	day := time.Now().Day()
	month := time.Now().Month()
	year := strconv.Itoa(time.Now().Year())

	dayStr, _ := createTwoDigits(day)
	montStr, _ := createTwoDigits(int(month))

	return fmt.Sprintf("%s_%s_%s", dayStr, montStr, year), nil
}
