package main

import (
	"log"
	"os"
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

	thumbnailData, err := createThumbnail(inputData)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("writing thumbnail to file..")
	err = os.WriteFile("thumbnail.jpg", thumbnailData, 0666)
	if err != nil {
		log.Fatal("error writing thumbnail", err)
	}
}
