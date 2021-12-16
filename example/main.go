package main

import (
	"log"

	"github.com/inuomi/go-makename"
)

func main() {
	nickname := makename.Generate()
	log.Println(nickname)
}
