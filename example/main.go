package main

import (
	"log"

	"github.com/cuipeiyu/go-makename"
)

func main() {
	nickname := makename.Generate()
	log.Println(nickname)
}
