package main

import (
	"log"
	"os/exec"
)

var str = [...]string{"hello", "haha"}

func main() {
	path, err := exec.LookPath("run-my-git-env.sh")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(path) // bin/ls

	log.Println(str)
}

// as util
func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
