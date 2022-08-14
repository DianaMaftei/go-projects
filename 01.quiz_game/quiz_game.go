package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

// https://github.com/gophercises/quiz

type problem struct {
	question string
	answer   string
}

func main() {

	fileName := flag.String("csv", "problems.csv", "Provide a csv in the form of 'question,answer'")
	shuffle := flag.Bool("shuffle", true, "Shuffle questions")
	timeLimit := flag.Int("limit", 10, "Time Limit for the quiz, in seconds")
	flag.Parse()

	quizData := readCSV(*fileName)
	if *shuffle {
		shuffleQuestions(quizData)
	}

	problems := buildProblemSet(quizData)

	fmt.Printf("There are %v questions. You have %v seconds to finish. Press enter to start the quiz when you are ready. \n", len(problems), *timeLimit)
	var start string
	fmt.Scanln(&start)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	quit := make(chan bool)
	go runQuiz(problems, timer.C, quit)
	<-quit
}

func runQuiz(problems []problem, time <-chan time.Time, quit chan bool) {
	var rightAnswers int

problemLoop:
	for _, problem := range problems {
		fmt.Println(problem.question)
		answerCh := make(chan string)

		go func() {
			var answer string
			fmt.Scanln(&answer)
			answerCh <- answer
		}()

		select {
		case <-time:
			break problemLoop
		case userAnswer := <-answerCh:
			if strings.EqualFold(strings.TrimSpace(problem.answer), strings.TrimSpace(userAnswer)) {
				rightAnswers++
			}
		}
	}
	displayFinalResult(len(problems), rightAnswers)
	quit <- true
}

func shuffleQuestions(quizData [][]string) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(quizData), func(i, j int) { quizData[i], quizData[j] = quizData[j], quizData[i] })
}

func readCSV(fileName string) [][]string {
	f, err := os.Open(fileName)

	handleError(err, fmt.Sprintf("Unable to open file %s", fileName))

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()

	handleError(err, fmt.Sprintf("Unable to read file %s", fileName))

	return data
}

func buildProblemSet(lines [][]string) []problem {
	problemSet := make([]problem, len(lines))

	for i, line := range lines {
		problemSet[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}

	return problemSet
}

func displayFinalResult(totalProblems int, rightAnswers int) {
	fmt.Printf("You got %v answer(s) right and %v answer(s) wrong out of %v total questions", rightAnswers, totalProblems-rightAnswers, totalProblems)
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
		os.Exit(1)
	}
}
