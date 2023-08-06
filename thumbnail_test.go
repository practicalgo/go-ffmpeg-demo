package main

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestCreateThumbnail(t *testing.T) {
	originalImage, err := os.ReadFile("testdata/book_cover.jpg")
	if err != nil {
		t.Fatal("error reading test file", err)
	}

	thumbnailImage, err := createThumbnail(context.Background(), originalImage, 30*time.Second)
	if err != nil {
		t.Fatal("error creating thumbnail", err)
	}
	// TODO assert the width and height of the image
	if len(thumbnailImage) == 0 || len(thumbnailImage) >= len(originalImage) {
		t.Fatal("thumbnail not created successfully")
	}

}
