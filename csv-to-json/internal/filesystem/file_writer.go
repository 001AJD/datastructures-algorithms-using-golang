// Package filesystem provides buffer io wrapper reader and writer
package filesystem

import (
	"bufio"
	"os"
)

type BufferedFileWriter struct {
	File           *os.File
	BufferedWriter *bufio.Writer
}

func NewBufferedFileWriter(path string) (*BufferedFileWriter, error) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	bufferedWriter := bufio.NewWriter(file)
	return &BufferedFileWriter{
		File:           file,
		BufferedWriter: bufferedWriter,
	}, nil
}

func (bw *BufferedFileWriter) Close() error {
	if err := bw.BufferedWriter.Flush(); err != nil {
		return err
	}
	if err := bw.File.Close(); err != nil {
		return err
	}
	return nil
}
