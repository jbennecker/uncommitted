package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

func main() {
	root, _ := os.Getwd()
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			if err := os.Chdir(fmt.Sprintf("%s/%s", root, file.Name())); err != nil {
				log.Fatal(err)
			}

			out, err := exec.Command("git", "status").Output()
			if err != nil {
				color.Blue("%s", file.Name())
				fmt.Println(err, "\n")
				continue
			}

			if bytes.Contains(out, []byte("nothing to commit, working directory clean")) != true {
				color.Blue("%s", file.Name())
				fmt.Printf("%s\n", out)
			}
		}
	}
}
