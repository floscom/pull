package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var path string

func main() {
	flag.StringVar(&path, "path", "./", "Specify the root path of your GIT projects.")
	flag.Parse()
	if path == "" {
		log.Fatal("PATH", "empty")
	}
	dir, err := os.Getwd()
	log.Println("MY PATH", dir)
	log.Println("PATH", path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			folder := fmt.Sprintf("%s/%s%s", dir, path, file.Name())
			log.Println("FOLDER", folder+"/.git")
			if _, err := os.Stat(folder + "/.git"); err == nil {
				log.Println("FOUND GIT", folder+"/.git")
				cmd := exec.Command("git", fmt.Sprintf("--git-dir=%s", folder+"/.git"), "pull")
				out, err := cmd.CombinedOutput()
				log.Println(string(out))
				if err != nil {
					log.Println("ERROR", err)
				}
			}
		}
	}
}
