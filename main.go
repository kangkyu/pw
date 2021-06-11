package main

import (
	"log"
	"flag"

	"github.com/sethvargo/go-password/password"
)

var length int

func init() {
	flag.IntVar(&length, "len", 9, "length of your password, should be longer than 5")
	flag.Parse()
}

func main() {
	if length < 5 {
		log.Fatalln("len flag should be greater or equal to 5")
	}
	res, err := password.Generate(length, 2, 2, false, false)
	if err != nil  {
		log.Fatal(err)
	}
	log.Printf(res)
}
