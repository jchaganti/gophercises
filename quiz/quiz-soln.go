package quiz

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type Problem struct {
	q string
	a int
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func parseLines(lines [][]string) []Problem {
	problems := make([]Problem, len(lines))
	for i, line := range lines {
		val, err := strconv.Atoi(line[1])
		handleError(err)
		problems[i] = Problem{
			q: line[0],
			a: val,
		}
	}
	return problems
}

func Main() {
	fileName := flag.String("fileName", "../quiz/problems.csv", "Enter a valid csv file")
	timeLimit := flag.Int("limit", 30, "Enter a time limit in seconds")
	flag.Parse()
	fmt.Printf("FileName entered is %s\n", *fileName)
	absPath, _ := filepath.Abs(*fileName)
	file, err := os.Open(absPath)
	handleError(err)
	// Close the file in the deferred manner(gets executed at the end of this function)
	defer file.Close()

	// Create a new csv reader
	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll()
	handleError(err)

	problems := parseLines(lines)
	score := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for _, problem := range problems {
		fmt.Printf("Add %+v\n", problem.q)
		answerCh := make(chan int)
		go func() {
			var sum int
			_, err = fmt.Scanf("%d\n", &sum)
			handleError(err)
			answerCh <- sum
		}()
		select {
		case <-timer.C:
			fmt.Printf("Time is up, your score is %d out of %d", score, len(problems))
			return
		case sum := <-answerCh:
			if problem.a == sum {
				score++
			}
		}

	}
	fmt.Printf("Your score is %d out of %d", score, len(problems))
}
