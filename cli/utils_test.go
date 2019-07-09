package main

import (
	"fmt"
	"testing"
)

func TestCurrentDir(t *testing.T) {
	currentPath := getCurrentPath()
	fmt.Printf("Current directory %s\n", currentPath)

	if currentPath == "" {
		t.Errorf("Current directory is empty")
	}
}

func TestDirExists(t *testing.T) {
	exists := checkDirExists("/usr/local/")
	if exists == false {
		t.Errorf("Dir exists check failed")
	}

	notExists := checkDirExists("/some/path/not/exists")
	if notExists == true {
		t.Errorf("Dir should not exists")
	}
}
