# Golang-Quiz-App

A simple quiz app with golang
This is a simple command-line quiz application written in Go. It reads quiz questions and answers from a CSV file, presents them to the user, and enforces a time limit for the quiz.

Features

1. Reads quiz questions from a CSV file.

2. Allows customization of the file name via a command-line flag.

3. Sets a timer to limit the quiz duration.

4. Automatically stops when the timer runs out.

5. Tracks the number of correct answers.

6. Displays the final score at the end of the quiz.

How It Works

1. The program reads quiz questions from a CSV file (default: quiz.csv).

2. It presents each question to the user one by one.

3. The user must enter their answer before the timer runs out.

4. The program keeps track of the number of correct answers.

5. When the timer expires, or all questions are answered, the program displays the final score.
