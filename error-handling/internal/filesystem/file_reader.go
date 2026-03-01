package filesystem

import (
	"bufio"
	"os"
)

type BufferedReader struct {
	File           *os.File
	BufferedReader *bufio.Reader
}

func NewBufferedFileReader(filePath string) (*BufferedReader, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0444)
	if err != nil {
		return nil, err
	}
	bufferedReader := bufio.NewReader(file)
	return &BufferedReader{
		File:           file,
		BufferedReader: bufferedReader,
	}, nil
}

func (b *BufferedReader) Close() error {
	if err := b.File.Close(); err != nil {
		return err
	}
	return nil
}
