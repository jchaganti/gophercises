package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func mymain() {
	fileName := flag.String("fileName", "default.csv", "Enter a valid csv file")
	flag.Parse()
	fmt.Printf("FileName entered is %s\n", *fileName)

	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	// Close the file in the deferred manner(gets executed at the end of this function)
	defer file.Close()
	// Create a new csv reader
	csvReader := csv.NewReader(file)
	var score, qnsCount int
	// Read the file till we reach EOF
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Add %+v\n", rec[0])
		result, err := strconv.Atoi(rec[1])
		if err != nil {
			log.Fatal("strconv error", err)
		}
		var sum int
		_, err = fmt.Scanf("%d\n", &sum)
		if err != nil {
			log.Fatal("Scanf error: ", err)
		}
		if result == sum {
			score++
		}
		qnsCount++
	}
	fmt.Printf("Your score is %d out of %d", score, qnsCount)
}
