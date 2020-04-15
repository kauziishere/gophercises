package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

const (
	defaultCSVFilename = "../resources/problems.csv"
	csvFlagHelpString  = "A csv filename in form of string to get question answers\nFormat for CSV file: question,answer"
	lFlagHelpString    = "Time limit to answer a question"
	defaultTimeLimit   = 300
)

type element struct {
	question string
	answer   string
}

var csvFilename *string
var timeLimit *int
var quiz []element

func readFileAndGenerateQuiz(csvFilename string) error {
	csvfile, err := os.Open(csvFilename)
	if nil != err {
		return err
	}

	csvReader := csv.NewReader(csvfile)

	for {
		line, err := csvReader.Read()
		if io.EOF == err {
			break
		}
		if nil != err {
			return err
		}

		varElement := element{question: line[0], answer: line[1]}

		quiz = append(quiz, varElement)
	}
	return nil
}

func conductQuiz(quiz []element, maxTime int) {
	var givenAnswer string
	score := 0
	totalQuestions := len(quiz)

	fetchAnswerChannel := make(chan string)

	fmt.Println("Starting Quiz:")
	for quesNo := 1; quesNo <= totalQuestions ; quesNo++ {
		fmt.Printf("Question no. %d: %s\nSolution: ", quesNo, quiz[quesNo-1].question)

		tick := time.Tick(time.Second * time.Duration(maxTime))

		go func() {
			answerReader := bufio.NewReader(os.Stdin)
			givenAnswer, _ = answerReader.ReadString('\n')
			fetchAnswerChannel <- givenAnswer
		}()

		fetchValueFailed := false

		for {
			select {
			case givenAnswer := <- fetchAnswerChannel:
				if quiz[quesNo-1].answer == strings.Replace(givenAnswer, "\n", "", 1) {
					score++
				}
			case <-tick:
				fmt.Printf("You took too long to answer")
				fetchValueFailed = true
			}
			break
		}

		if true == fetchValueFailed {
			break
		}
	}

	fmt.Printf("Your final score is %d out of %d\n", score, totalQuestions)
}

func init() {
	csvFilename  = flag.String("csv", defaultCSVFilename, csvFlagHelpString)
	timeLimit = flag.Int("l", defaultTimeLimit, lFlagHelpString)
	flag.Parse()
}

func main() {
	var err error

	err = readFileAndGenerateQuiz(*csvFilename)
	if nil != err {
		log.Panic(err.Error())
	}

	conductQuiz(quiz, *timeLimit)
}
