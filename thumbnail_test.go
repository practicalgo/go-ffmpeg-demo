package main

import (
	"os"
	"testing"
)

func TestCreateThumbnail(t *testing.T) {
	originalImage, err := os.ReadFile("testdata/book_cover.jpg")
	if err != nil {
		t.Fatal("error reading test file", err)
	}

	thumbnailImage, err := createThumbnail(originalImage)
	if err != nil {
		t.Fatal("error creating thumbnail", err)
	}
	// TODO assert the width and height of the image
	if len(thumbnailImage) == 0 || len(thumbnailImage) >= len(originalImage) {
		t.Fatal("thumbnail not created successfully")
	}

}
