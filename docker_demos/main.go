package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "test1":
		separateNamespace()
	case "test2":
		printFile()
	default:
		panic("no fucking this command!")
	}
}

func separateNamespace() {
	fmt.Printf("Running: %v\n", os.Args[2:])

	//cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd := exec.Command("sh")

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func printFile() {

	ioutil.ReadFile("")

}
