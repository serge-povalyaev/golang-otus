package main

import (
	"errors"
	"io"
	"os"

	pb "github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrFileNotFound          = errors.New("file not found")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	var sourceFile *os.File
	sourceFile, err := os.Open(fromPath)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrFileNotFound
		}
		return err
	}
	defer sourceFile.Close()

	fileInfo, err := sourceFile.Stat()
	if err != nil {
		return err
	}

	fileSize := fileInfo.Size()
	if fileSize == 0 || fileInfo.IsDir() {
		return ErrUnsupportedFile
	}

	if offset > fileSize {
		return ErrOffsetExceedsFileSize
	}

	targetFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	if limit == 0 {
		limit = fileSize - offset
	}

	n := offset + limit

	bar := pb.StartNew(getBarSize(fileSize, offset, limit))

	buf := make([]byte, 1)
	for offset < n {
		_, err := sourceFile.ReadAt(buf, offset)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		_, err = targetFile.Write(buf)
		if err != nil {
			return err
		}

		bar.Increment()

		offset++
	}

	bar.Finish()

	return nil
}

func getBarSize(fileSize, offset, limit int64) int {
	if fileSize > offset+limit {
		return int(limit)
	}

	return int(fileSize - offset)
}
