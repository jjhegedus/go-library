package main

import (
	"log"
	"os"

	entities "github.com/jjhegedus/go-library/entities"
)

func main() {
	l := log.New(os.Stdout, "github.com/jjhegedus/go-library", log.LstdFlags)
	l.Print("main:begin")

	entities.Entities()
}
