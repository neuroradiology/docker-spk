package main

// Embed the sandstorm-pkgdef.capnp template in the docker-spk executable.

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

const OutFile = "template.gen.go"

func chkfatal(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	bytes, err := ioutil.ReadFile("sandstorm-pkgdef.capnp.template")
	chkfatal(err)
	file, err := os.Create(OutFile)
	chkfatal(err)
	_, err = fmt.Fprintf(file,
		`package main

// This file was auto-generated by ./internal/embed-template.go.
// DO NOT EDIT.

import "text/template"

var PkgDefTmplate = template.Must(template.New("pkgdef").Parse(%q))
`, bytes)
	chkfatal(err)
	chkfatal(file.Close())
	chkfatal(exec.Command("gofmt", "-w", OutFile).Run())
}