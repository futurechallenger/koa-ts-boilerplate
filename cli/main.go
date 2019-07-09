package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"runtime"
)

func main() {
	// if checkSystem() == true {
	// 	// Handle windows stuff
	// } else {
	// 	// Handle non-windows stuff
	// }

	if checkGit() == false {
		fmt.Println("Git is not installed!")
		return
	}

	currentPath := getCurrentPath()
	if checkDirExists(currentPath) == true {
		fmt.Println("Directory already exists!")
		return
	}

	// TODO: 1. create app directory. 2. copy files into the directory
	destDir := fmt.Sprintf("%s//%s", currentPath, "my-app")

	createdOK := createDir(destDir)
	if createdOK == false {
		fmt.Println("Create dir error")
		return
	}

	// TODO: copy files,
	// 1. No git files
	// 2. No cli files
	// 3. No readme.md (optional)
	// 4. No node_modules
	copyDir()
}
