package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	//"path"

	// CaltechLibrary package
	"github.com/caltechlibrary/lg2md"
)

func main() {
	//appName := path.Base(os.Args[0])
	flag.Parse()
	args := flag.Args()

	inputFName := args[0]
	src, err := ioutil.ReadFile(inputFName)
	if err != nil {
		log.Fatal(err)
	}
	lg, err := lg2md.Decode(lg2md.Clean(src))
	if err != nil {
		log.Fatal(err)
	}
	s, err := lg.ToJSON()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(os.Stdout, "%s", s)
}
