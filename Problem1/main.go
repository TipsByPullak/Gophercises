package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
)

//Struct to track the stats of the quiz
type quizStat struct {
	noOfCorrect int
	noOfTotal   int
}

func main() {
	//Read the flag and parse it
	filePtr := flag.String("filename", "problems.csv", "Used to provide the filename to the executable")
	flag.Parse()

	//Testing the flag parser-
	//fmt.Println(*filePtr)

	//Create the iterator to read each question and answer
	problemReader := csvReader(*filePtr)

	//Start the Quiz and print stats at the end
	stat := startQuiz(problemReader)
	fmt.Printf("\n-----QUIZ END-----.\nNo. of correct answers= %v\nNo. of Questions Attempted= %v\n\n", stat.noOfCorrect, stat.noOfTotal)
}

//Function to read the CSV file
func csvReader(filename string) *csv.Reader {
	file, err := os.Open(filename) //Open the file
	//If file opening fails
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(2)
	}

	return csv.NewReader(file)
}

func startQuiz(reader *csv.Reader) quizStat {
	stat := quizStat{} //the struct to be returned

	//Quiz only starts once the user presses enter
	fmt.Println("Press Enter to start the quiz..")
	fmt.Scanln()

	var answer string
	for {
		question, err := reader.Read() //Read the next question
		if err == io.EOF {             //If all questions have been read
			break
		}
		if err != nil { //If there's an error in reading the file
			fmt.Println("Error:", err)
		}

		(stat.noOfTotal)++                                      //Total number of questions attempted is increased
		fmt.Printf("Q%v: %v\nA: ", stat.noOfTotal, question[0]) //Display the question
		fmt.Scanln(&answer)                                     //Read the answer //Implement string trimming
		if answer == question[1] {
			(stat.noOfCorrect)++ //Increase number of correct answers if answered correctly
		}
	}

	return stat
}
