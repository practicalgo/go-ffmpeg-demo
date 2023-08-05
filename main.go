package main

import (
	"context"
	"log"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("expected an image file path as argument")
	}
	inputImagePath := os.Args[1]
	inputData, err := os.ReadFile(inputImagePath)
	if err != nil {
		log.Fatal("error reading input", err)
	}

	thumbnailData, err := createThumbnail(
		context.Background(),
		inputData,      // image bytes to create thumbnails
		10*time.Second, // command timeout
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("writing thumbnail to file..")
	err = os.WriteFile("thumbnail.jpg", thumbnailData, 0666)
	if err != nil {
		log.Fatal("error writing thumbnail", err)
	}
}
