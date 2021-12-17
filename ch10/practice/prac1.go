package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"
)

var targetType = flag.String("imgType", "png", "target image type")

func main() {
	flag.Parse()

	if err := toTargetType(os.Stdin, os.Stdout, *targetType); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg:%v\n", err)
		os.Exit(1)
	}
}

func toTargetType(in io.Reader, out io.Writer, targetType string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	targetType = strings.ToLower(targetType)
	if targetType == "jpeg" {
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	} else if targetType == "png" {
		return png.Encode(out, img)
	} else if targetType == "gif" {
		return gif.Encode(out, img, &gif.Options{})
	} else {
		return nil
	}
}
