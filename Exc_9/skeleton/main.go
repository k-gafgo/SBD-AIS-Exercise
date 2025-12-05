package main

import (
	"bytes"
	"exc9/mapred"
	"fmt"
	"log"
	"os"
	"strings"
)

// Main function
func main() {
	// todo read file
	dat, err := os.ReadFile("res/meditations.txt")
	if err != nil {
		log.Printf("Error reading file: %v", err.Error())
	}
	dat_str := string(dat[:])
	text := strings.Fields(dat_str)

	// todo run your mapreduce algorithm
	var mr mapred.MapReduce
	results := mr.Run(text)
	// todo print your result to stdout
	b := new(bytes.Buffer)
	for key, value := range results {
		fmt.Fprintf(b, "%s:\"%s\"\n", key, value)
	}

	_, err = os.Stdout.WriteString(b.String())
	if err != nil {
		log.Printf("Error writing file: %v", err.Error())
	}
}
