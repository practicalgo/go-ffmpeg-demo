package main

import (
	"context"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"time"
)

func createThumbnail(
	cmdCtx context.Context,
	inputData []byte,
	cmdTimeout time.Duration,
) ([]byte, error) {

	var err error

	ch := make(chan error, 1)
	// TODO image id
	tmpDir, err := os.MkdirTemp("", "thumbnail-")
	if err != nil {
		log.Println("error creating temporary directory", err)
		return nil, err
	}
	defer func() {
		err := os.RemoveAll(tmpDir)
		if err != nil {
			log.Println("error cleaning up", tmpDir)
		}
	}()

	inputFile, err := os.CreateTemp(tmpDir, "input-image")
	if err != nil {
		return nil, err
	}
	_, err = io.WriteString(inputFile, string(inputData))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(cmdCtx, cmdTimeout)
	defer cancel()

	// ffmpeg will work with stdin and stdout
	// we set those up next
	outputFile := path.Join(tmpDir, "thumbnail-image-id")
	cmd := exec.CommandContext(
		ctx,
		"ffmpeg", "-i", inputFile.Name(), "-vf", "scale='iw/2:ih/2'", "-f", "image2", outputFile,
	)

	go func() {

		log.Println("creating thumbnail of a specified image using ffmpeg..")
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Println("error creating thumbnail", string(output))
		}

		ch <- err
	}()

	select {
	case <-ctx.Done():
		err = ctx.Err()
	case err = <-ch:
	}

	if err != nil {
		return nil, err
	}

	thumbnailBytes, err := os.ReadFile(outputFile)
	if err != nil {
		log.Println("error reading from output file", outputFile)
		return nil, err
	}

	return thumbnailBytes, err
}
