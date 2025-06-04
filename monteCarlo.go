package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func main() {
	var diceSides, diceRolls, minimumScore int
	var tempTotal, tempRolls int
	var customValues string

	fmt.Println("This program will roll a set of dice 10.000 times.")
	fmt.Println("You can set how many sides are on your dice, and how many dice are in the set.")
	fmt.Println("You can even set custom values for the dice sides.")

	fmt.Println("How many sides on your dice?")
	fmt.Scanln(&diceSides)

	customDiceSides := make([]int, diceSides)

	fmt.Println("Do you want to use custom values for your dice sides? (Y/n)")
	fmt.Scanln(&customValues)

	if customValues == "Y" || customValues == "y" || customValues == "" {
		fmt.Println("Please enter the values for each side one by one")
		for i := 0; i < diceSides; i++ {
			fmt.Println("Enter value for side", i+1)
			fmt.Scanln(&customDiceSides[i])
		}
	} else {
		for i := 0; i < diceSides; i++ {
			customDiceSides[i] = i + 1
		}
	}

	fmt.Println("How many dice are you rolling?")
	fmt.Scanln(&diceRolls)

	fmt.Println("What's the minimum roll you're aiming for?")
	fmt.Scanln(&minimumScore)

	// create the slice storing the rolls
	// To do: Update this for customDiceSides
	monteCarlo := make([]int, (slices.Max(customDiceSides)*diceRolls)+1)

	// create a slice to store the rolls needed for the minimum target score
	rollsNeeded := make([]int, diceRolls+slices.Max(customDiceSides))

	// The Monte Carlo magic happens here
	for i := 0; i < 10000; i++ {
		tempTotal, tempRolls = rollTotal(diceSides, diceRolls, minimumScore, customDiceSides)
		monteCarlo[tempTotal] += 1
		rollsNeeded[tempRolls] += 1
	}

	// Print the possible outcomes
	for i, n := range monteCarlo {
		// Skip the impossible outcomes
		if n >= diceRolls {
			fmt.Println("Amount of", i, ":", n)
		}
	}

	// Print the rolls needed for the minimum score
	for i, n := range rollsNeeded {
		if i == 0 {
			fmt.Println("In", n, "cases, the minimum score wasn't rolled.")
			var percent int
			percent = (10000 - n) / 100
			fmt.Println("In other words, the chance to roll the minimum or higher is", percent, "%")
		}
		if i != 0 && n != 0 {
			fmt.Println("Amount of rolls required for", minimumScore, ":", i, "in", n, "cases.")
		}
	}
}

// Rolls some dice.
func rollTotal(diceSides, diceRolls, minimumScore int, customDiceSides []int) (int, int) {
	total := 0
	minrolls := 0
	for i := 0; i < diceRolls; i++ {
		total += customDiceSides[rand.Intn(diceSides)]
		if total >= minimumScore && minrolls == 0 {
			minrolls = i + 1
		}
	}
	return total, minrolls
}
