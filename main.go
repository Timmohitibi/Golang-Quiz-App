package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func problemPuller(fileName string)([]problem, error){
	// read all problems from quiz.csv
	// open file
	if fObj, err := os.Open(fileName); err == nil{
		//create a new reader
		csvR := csv.NewReader(fObj)
		// will need to read file
		if cLines, err := csvR.ReadAll(); err == nil{
			return parseProblem(cLines), nil
		}else{
			return nil, fmt.Errorf("error reading csv file: %s", err.Error())

		}

		}else{
		return nil, fmt.Errorf("error opening file: %s", err.Error())
		}
	}

func main(){
	//1. input name of file
	fName := flag.String("f", "quiz.csv", "path of the file to read")
	//2. set duration of timer
	timer :=  flag.Int("t", 30, "timer for the quiz")
	flag.Parse()
	//3. pull problems from file-- problem piller function
	problems, err := problemPuller(*fName)
	//4. Handle error
	if err != nil{
		exit(fmt.Sprintf("Error reading file: %s", err.Error()))
	}
	//5. create variable to check correct answers
	correctAns := 0
	//6. initialize the timer
	tObj := time.NewTimer(time.Duration(*timer) * time.Second)
	ansC := make(chan string)
	//7. loop through problems 
	problemLoop:
	for i, p := range problems{
		var answer string
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		// start a goroutine to read the answer
		go func(){
			fmt.Scanf("%s", &answer)
			ansC <- answer
		}()
		// start a goroutine to read the answer
		select {
		case <-tObj.C:
			fmt.Println()
			break problemLoop
		case iAns := <- ansC:
			if iAns == p.a{
				correctAns++
			}
			if i == len(problems)-1{
				close(ansC)
			}
		}
	}


	//8. check if answer is correct
	fmt.Printf("Your result is %d out of %d/n", correctAns, len(problems))
	fmt.Printf("press enter to exit")
	<- ansC
	//9. if correct, add to score
	//10. if incorrect, add to score
}

func parseProblem(lines [][]string) []problem{
	// go ove r the lines and pass them to problem struct
	r := make([]problem, len(lines))
	for i := 0; i<len(lines); i++{
		r[i] = problem{
			q: lines[i][0],
			a: lines[i][1],
		}
	}
	return r
	// return a slice of problem struct
}


type problem struct{
	q string
	a string
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}