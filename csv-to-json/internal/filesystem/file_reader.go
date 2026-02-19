// Package filesystem provides buffer io wrapper reader and writer
package filesystem

import (
	"bufio"
	"os"
)

type BufferedFileReader struct {
	File           *os.File
	BufferedReader *bufio.Reader
}

func NewBufferedFileReader(path string) (*BufferedFileReader, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0444)
	if err != nil {
		return nil, err
	}
	bufferedReader := bufio.NewReader(file)
	return &BufferedFileReader{
		File:           file,
		BufferedReader: bufferedReader,
	}, nil
}

func (br *BufferedFileReader) Close() error {
	if err := br.File.Close(); err != nil {
		return err
	}
	return nil
}
