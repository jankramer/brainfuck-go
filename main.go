package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	path := flag.String("f", "", "path to brainfuck program to run")
	flag.Parse()

	if *path == "" {
		flag.Usage()
		os.Exit(1)
	}

	program, err := ioutil.ReadFile(*path)
	if err != nil {
		log.Fatalln("Failed to read program:", err)
	}

	if err := Run(program, os.Stdin, os.Stdout); err != nil {
		log.Fatalln("Failed to run program:", err)
	}
}
