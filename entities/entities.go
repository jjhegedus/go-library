package entities

import (
	"log"
	"os"
)

func Entities() {
	l := log.New(os.Stdout, "github.com/jjhegedus/go-library/entities", log.LstdFlags)
	l.Print("Entities:begin")
}
