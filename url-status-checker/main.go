package main

import (
	"fmt"
)

func main() {
	fmt.Println("URL status checker initiating...")
	ProcessFile("./data/top-100.csv")
}
