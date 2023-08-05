package main

import (
	"bytes"
	"context"
	"io"
	"log"
	"os/exec"
	"time"
)

func createThumbnail(
	cmdCtx context.Context,
	inputData []byte,
	cmdTimeout time.Duration,
) ([]byte, error) {

	var cmdOut, cmdErr bytes.Buffer
	var err error

	ch := make(chan error, 1)

	ctx, cancel := context.WithTimeout(cmdCtx, cmdTimeout)
	defer cancel()

	// ffmpeg will work with stdin and stdout
	// we set those up next
	cmd := exec.CommandContext(
		ctx,
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

	go func() {

		// we cannot use CombinedOutput() here as it returns both
		// stdout and stderr
		// stdout will contain our transformed image when successful
		// we also want to report the entire error when there is a problem
		// so, we have to separate out stdout and stderr explicitly
		//out, err := cmd.CombinedOutput()

		log.Println("creating thumbnail of a specified image using ffmpeg..")
		err = cmd.Run()
		if err != nil {
			log.Println(string(cmdErr.Bytes()))
		}
		ch <- err
	}()

	select {
	case <-ctx.Done():
		err = ctx.Err()
	case err = <-ch:
	}

	return cmdOut.Bytes(), err
}
