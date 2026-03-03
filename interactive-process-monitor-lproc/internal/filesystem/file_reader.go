package filesystem

import (
	"bufio"
	"fmt"
	"os"
)

type BufferedFileReader struct {
	File           *os.File
	BufferedReader *bufio.Reader
}

func NewBufferedFileReader(filepath string) (*BufferedFileReader, error) {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0444)
	if err != nil {
		return nil, err
	}
	bufferedReader := bufio.NewReader(file)
	return &BufferedFileReader{
		File:           file,
		BufferedReader: bufferedReader,
	}, nil
}

func (b *BufferedFileReader) Close() {
	err := b.File.Close()
	if err != nil {
		fmt.Println("Error occured when closing file handler")
	}
}
