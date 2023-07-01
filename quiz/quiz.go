package quiz

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

type MyProblem struct {
	q string
	a int
}

func myParseLines(lines [][]string) (problems []MyProblem) {
	problems = make([]MyProblem, len(lines))
	for i, line := range lines {
		a, err := strconv.Atoi(line[1])
		handleError(err)
		problems[i] = MyProblem{
			q: line[0],
			a: a,
		}
	}
	return
}

func QuizMain() {
	fileName := flag.String("fileName", "../quiz/problems.csv", "Enter a valid csv file")
	timeLimit := flag.Int("limit", 10, "Time limit for the quiz")
	flag.Parse()
	fmt.Printf("FileName entered is %s, time limit is %d\n", *fileName, *timeLimit)

	file, err := os.Open(*fileName)
	handleError(err)
	// Close the file in the deferred manner(gets executed at the end of this function)
	defer file.Close()
	// Create a new csv reader
	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll()
	handleError(err)
	problems := myParseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	var score int
	for _, problem := range problems {
		fmt.Println(problem.q)
		answerCh := make(chan int)
		go func() {
			var sum int
			_, err = fmt.Scanf("%d\n", &sum)
			handleError(err)
			answerCh <- sum
		}()
		select {
		case <-timer.C:
			fmt.Printf("Time up. Your score is %+v", score)
			return
		case sum := <-answerCh:
			if sum == problem.a {
				score++
			}
		}
	}
	fmt.Printf("Your score is %d out of %d", score, len(problems))
}
