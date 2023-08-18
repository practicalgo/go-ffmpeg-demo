# Demo of using `ffmpeg` from Go

[![Build and Test](https://github.com/practicalgo/go-ffmpeg-demo/actions/workflows/main.yml/badge.svg?branch=main)](https://github.com/practicalgo/go-ffmpeg-demo/actions/workflows/main.yml)

This demo specifically shows how we can supply input and obtain output via using 
input and output pipes to `ffmpeg`. 

This demo has been verified to be able to create a thumbnail of an image on MacOS, 
Linux and Windows Server 2019 and 2022. See the 
[GitHub workflow runs](https://github.com/practicalgo/go-ffmpeg-demo/actions/workflows/main.yml) 
if you want to learn more.

## Implementation

See [thumbnail.go](./thumbnail.go) for the execution of `ffmpeg`. 

## Improvements

- Use context to ensure that we don't wait for ever for `ffmpeg` to complete. See [this commit](https://github.com/practicalgo/go-ffmpeg-demo/commit/efc70f0514d9cc02f896f354b4dd9da1a2afac9a) for one way to do so. I removed the change to simply this demo itself.
- Update test to verify the dimensions of the thumbnail to be created

