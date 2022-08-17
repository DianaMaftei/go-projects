package main

import (
	"fmt"
	"time"
)

func main() {
	animal := "cat"
	makeNoise(animal)

	today := time.Now().Weekday()
	checkWeekend(today)

	checkAgeForDrinking(17)

}

func makeNoise(animal string) {
	switch animal {
	case "cat":
		fmt.Println("Meow!")
	case "dog":
		fmt.Println("Woof, Woof!")
	case "cow":
		fmt.Println("Moo!")
	default:
		fmt.Println("Animal has not been observed")
	}
}

func checkWeekend(weekday time.Weekday) {
	switch weekday {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the Weekend! Woohoooo!")
	default:
		fmt.Println("Sadly, it's not the weekend!")
	}
}

func checkAgeForDrinking(age int) {
	switch {
	case age < 18:
		fmt.Println("You can't drink in Europe or the US")
	case age < 21:
		fmt.Println("You can drink in Europe, but you can't drink in the US")
	default:
		fmt.Println("You can drink both in Europe, and in the US")
	}
}
