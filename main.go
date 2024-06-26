package main

import (
	"log"

	"github.com/NabinGrz/SocialMediaApi/src/router"
)

func main() {

	r := router.Router()
	if err := r.Run(":3000"); err != nil {
		log.Panicf("error: %s", err)
	}

}
