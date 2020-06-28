// Copyright (c) 2020 ykcab
// You should have received a copy of a LICENSE which came along
// with this package distribution

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "%s\n", "wrong or no arguments value provided")
		return
	}
	if err := newGoFile(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", "failed to create "+os.Args[1]+" project. error: %s\n", err)
		return
	}

}

func newGoFile(s string) error {
	const fileExtension = ".go"

	tplt := []byte("package " + s)

	// create the folder
	folder := fmt.Sprintf("%s", os.Getenv("GOPATH")) + "/src/" + s + "/"
	if err := newDir(folder); err == nil {
		file := folder + s + fileExtension

		// check if the folder and file already exist
		if _, err := os.Stat(file); err == nil {
			fmt.Fprintf(os.Stderr, "%s\n", "cannot create "+s+fileExtension+". file and folder already exists. try another name")
			return err
		}
		_, err := os.Create(file)
		if err != nil {
			return err
		}
		if err := ioutil.WriteFile(file, tplt, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return err
		}

	} else {
		log.Fatalf("could not create project directory folder. error: %s", err)
	}
	return nil
}

func newDir(s string) error {
	_, err := os.Stat(s)

	// check if this folder exsits or not
	// create it otherwise
	if os.IsNotExist(err) {
		err := os.MkdirAll(s, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
