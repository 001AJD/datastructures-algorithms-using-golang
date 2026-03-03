package processor

import (
	"encoding/csv"
	"fmt"
	"io"
	"lproc/internal/filesystem"
)

func Init() {
	bufferedFileReader, err := filesystem.NewBufferedFileReader("./data/processes.csv")
	if err != nil {
		fmt.Println("error occurred when trying to read the file", err)
	}
	csvReader := csv.NewReader(bufferedFileReader.File)

	for {
		recorder, err := csvReader.Read()
		if err == io.EOF {
			fmt.Println("End of file reached")
			break
		}
		if err != nil {
			fmt.Println("Error reading a record from the csv file, skipping and moving to the next record")
		}
		fmt.Println(recorder)
	}
}
