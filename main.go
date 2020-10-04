package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Template io.Reader
	Out io.Writer
}

// We are just going to start with basic flags but then probably update to an actual command framework like cobra
func main(){
	cfg := setup()

	// Just output the template.
	template, _ := ioutil.ReadAll(cfg.Template)
	fmt.Println(string(template))
}


func setup() Config {
	templateFile := flag.String("template", "", "file to change from json to cdk")
	outFile := flag.String("out", "cdk.ts", "output file with transformed typescript")
	flag.Parse()

	if *templateFile == "" {
		log.Fatal("Need to specify a file via -template")
	}

	// for now just put it all into a byte array, maybe clean this up but I don't like the idea
	// of having a file hanging open, when memory usage isn't too bad at this point.
	tmp, err := ioutil.ReadFile(*templateFile)

	if err != nil {
		log.Fatalln("Could not read file", *templateFile, "please check permissions/existence", err)
	}

	fd, err:= os.Create(*outFile)

	if err != nil {
		log.Fatalln("Could not create file to write to at", *outFile, "please check if this file should be there.")
	}
	// TODO - what's a good way to close the file output..
	return Config{ Template: bytes.NewBuffer(tmp), Out: fd}
}