package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func readOut(reader io.ReadCloser, f *os.File) (bool, error) {
	buf := make([]byte, 16)
	n, err := reader.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	if n > 0 {
		_, err := f.Write(buf[:n])
		if err != nil {
			log.Fatalf("Unable to write to file: %s", err)
		}
	}
	return n == 0, nil
}

func main() {
	cmd := exec.Command("unbuffer", "python3", "target.py")
	targetStdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	targetStderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	f, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0655)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	for {
		doneStdin, err := readOut(targetStdout, f)
		if err != nil {
			log.Fatal(err)
		}
		doneStdout, err := readOut(targetStderr, f)
		if err != nil {
			log.Fatal(err)
		}

		if doneStdin && doneStdout {
			break
		}
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}
