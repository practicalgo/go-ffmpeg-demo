package main

import (
	"bytes"
	"io"
	"log"
	"os/exec"
)

func createThumbnail(inputData []byte) ([]byte, error) {

	var cmdOut, cmdErr bytes.Buffer
	var err error

	// ffmpeg will work with stdin and stdout
	// we set those up next
	cmd := exec.Command(
		"ffmpeg", "-i", "pipe:0", "-vf", "scale='iw/2:ih/2'", "-f", "image2", "pipe:1",
	)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	cmd.Stdout = &cmdOut
	cmd.Stderr = &cmdErr

	go func() {
		defer stdin.Close()
		_, err = io.WriteString(stdin, string(inputData))
	}()

	// we cannot use CombinedOutput() here as it returns both
	// stdout and stderr
	// stdout will contain our transformed image when successful
	// we also want to report the entire error when there is a problem
	// so, we have to separate out stdout and stderr explicitly

	err = cmd.Run()
	if err != nil {
		log.Println(string(cmdErr.Bytes()))
	}

	return cmdOut.Bytes(), err
}
