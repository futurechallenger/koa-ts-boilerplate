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

// createDir, with destination directory.
// Return true if success, or return false.
func createDir(destDir string) bool {
	err := os.Mkdir(destDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Error in creating destination directory")
		return false
	}

	return true
}

// copyFile copy a file from source to destination
func copyFile(src string, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

// copyDir copy source dir files to dest dir recursively
func copyDir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = copyDir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = copyFile(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}
