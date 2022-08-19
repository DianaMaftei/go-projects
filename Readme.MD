# A collection of projects in GO

## 00.learn_go

Contains code examples based on what I learned in the "Go: The Complete Developer's
Guide" course by Stephen Grider, on Udemy (lessons 00 to 13), as well as from the "https://gobyexample.com/" website, for things that weren't covered in the course (14-onwards).

## 01.quiz_game - (from https://github.com/gophercises/quiz)
A program that quizzes a user with questions from a file. 
- uses flags to pass in csv file name, time limit & if questions should be shuffled
- reads from csv file
- (optionally) shuffles the questions
- prompts user for answer from keyboard
- runs a timer that stops the quiz when it runs out
- displays quiz results to user
- 
## 02.url_shortener - (from https://github.com/gophercises/urlshort)
A program that acts as a url shortener, by looking at the path of an incoming HTTP request and redirecting it if it finds a matching url
- creates an HTTP server and an HTTP Handler to redirect incoming requests
- parses yaml and json files for path to match and url to redirect to
- populates a DB with an SQL script that creates a table and inserts paths and urls
- reads paths and urls from DB
