package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func checkSystem() bool {
	fmt.Printf("OS: %s\n", runtime.GOOS)
	return runtime.GOOS == "windows"
}

func checkGit() bool {
	output, err := exec.Command("git", "--version").CombinedOutput()
	fmt.Printf("Git installed %s", output)
	return err == nil
}

func getCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	return dir
}

// Return true is exists
func checkDirExists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}

	return false
}

func main() {
	// if checkSystem() == true {
	// 	// Handle windows stuff
	// } else {
	// 	// Handle non-windows stuff
	// }

	if checkGit() == false {
		fmt.Println("Git is not installed!\n")
		return
	}

	currentPath := getCurrentPath()
	if checkDirExists(currentPath) == true {
		fmt.Println("Directory already exists!")
		return
	}

	// TODO: 1. create app directory. 2. copy files into the directory
}
