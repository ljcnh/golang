package main

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var sourceFile = flag.String("sourceFile", "", "source file")
var targetFile = flag.String("targetFile", "", "target file")

//只能解压zip和tar

func main() {
	flag.Parse()
	if *sourceFile == "" {
		fmt.Println("sourceFile nil")
		return
	}
	if ExistDir(*sourceFile) {
		fmt.Println("file no exit")
		return
	}
	_, fileName := filepath.Split(*sourceFile)
	fileType := strings.Split(fileName, ".")
	l := len(fileType)
	if l == 0 {
		fmt.Println("file error,only tar and zip")
		return
	}
	if fileType[l-1] == "zip" {
		unZip(*sourceFile, *targetFile)
	} else if fileType[l-1] == "tar" || (l > 1 && fileType[l-2] == "tar") {
		unTar(*sourceFile, *targetFile)
	} else {
		fmt.Println("only tar and zip")
	}
}

func unTar(src, dst string) error {
	fr, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fr.Close()

	gr, err := gzip.NewReader(fr)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		case hdr == nil:
			continue
		}
		dstFileDir := filepath.Join(dst, hdr.Name)
		switch hdr.Typeflag {
		case tar.TypeDir:
			if b := ExistDir(dstFileDir); !b {
				if err := os.MkdirAll(dstFileDir, 0775); err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			}
		case tar.TypeReg:
			file, err := os.OpenFile(dstFileDir, os.O_CREATE|os.O_RDWR, os.FileMode(hdr.Mode))
			if err != nil {
				return err
			}
			_, err = io.Copy(file, tr)
			if err != nil {
				return err
			}
			file.Close()
		}
	}
	return nil
}

func unZip(src, dst string) error {
	zr, err := zip.OpenReader(src)
	defer zr.Close()
	if err != nil {
		return err
	}
	if dst != "" {
		if err := os.MkdirAll(dst, 0755); err != nil {
			return err
		}
	}
	for _, file := range zr.File {
		path := filepath.Join(dst, file.Name)
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(path, file.Mode()); err != nil {
				return err
			}
			continue
		}
		fr, err := file.Open()
		if err != nil {
			fr.Close()
			return err
		}
		fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, file.Mode())
		if err != nil {
			fr.Close()
			fw.Close()
			return err
		}
		_, err = io.Copy(fw, fr)
		if err != nil {
			fr.Close()
			fw.Close()
			return err
		}
		fw.Close()
		fr.Close()
	}
	return nil
}

func ExistDir(dirname string) bool {
	fi, err := os.Stat(dirname)
	return (err == nil || os.IsExist(err)) && fi.IsDir()
}
