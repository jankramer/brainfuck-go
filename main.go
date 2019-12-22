package main

import (
	"flag"
	"fmt"
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

	log.Println("Reading from", *path)

	program, err := ioutil.ReadFile(*path)
	if err != nil {
		log.Fatalln("Failed to read program:", err)
	}

	output, err := Run(program)

	if err != nil {
		log.Fatalln("Failed to run program:", err)
	}

	fmt.Println(string(output))
}
