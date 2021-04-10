package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main()  {
	csvFileName := flag.String("csv", "problems.csv",
		"A csv file in the format of 'question, answer'")
	timeLimit := flag.Int("limit", 30, "the time limit in seconds")
	flag.Parse()

	file, err :=os.Open(*csvFileName)
	if err!=nil {
		exit(fmt.Sprintf("failed to open the svc file: %s\n", *csvFileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err!=nil {
		exit("Failed to parse")
	}
	//fmt.Println(lines)
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i, p := range problems{
		fmt.Printf("Problem number %d: %s\n", i + 1, p.q )
		answerChan := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)

			answerChan<- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d \n", correct, len(problems))
			return
		case answer := <-answerChan:
			if answer == p.a{
				correct++
			}
		}
	}

}

func parseLines(lines[][] string) []problem {
	ret := make([] problem, len(lines))

	for i, line := range lines {
		ret[i] = problem {
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct{
	q string
	a string
}

func exit(msg string)  {
	fmt.Println(msg)
	os.Exit(1)
}

