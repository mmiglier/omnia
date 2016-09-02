package main

//go:generate go-bindata -pkg tool -o builtin/tool/data.go builtin/tool/data/...

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mmiglier/omnia/command"
)

const omniaDir = ".omnia"
const omniafileName = "omnia.server.yml"

func main() {
	os.Exit(realMain())
}

func realMain() int {

	var version bool
	var compile bool
	var deploy bool

	flag.BoolVar(&version, "version", false, "Print version information and quit")
	flag.BoolVar(&compile, "compile", false, "Create omnia infrastructure code")
	flag.BoolVar(&deploy, "deploy", false, "Deploy omnia infrastructure")

	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		return 1
	}

	if version {
		fmt.Printf("Omnia version %s, build %s\n", Version, Build)
		return 0
	}

	if compile {
		log.Println("Compiling...")
		if err := command.Compile(omniafileName, omniaDir); err != nil {
			log.Fatal("Error: ", err)
		}
		log.Println("Done!")
	}

	if deploy {
		log.Println("Deploying...")
		if err := command.Deploy(omniafileName, omniaDir); err != nil {
			log.Fatal("Error: ", err)
		}
		log.Println("Done!")
	}

	return 0
}
