package main

import (
	"fmt"
	"math/rand"
)

var (
	level          = 1
	nums     int32 = 0
	tempNums int32 = 0
	count          = 0
)

func main() {

	var m map[int]string = map[int]string{1: "Easy", 2: "Medium", 3: "Hard"}

	welcome := "Welcome to the Number Guessing Game!\nI'm thinking of a number between 1 and 100.\nYou have 5 chances to guess the correct number."
	fmt.Println(welcome)
	fmt.Println("")

	for {
		nums = rand.Int31n(100) + 1 // 1-100
		fmt.Println("Please select the difficulty level:")
		fmt.Println("1. Easy (10 chances)\n2. Medium (5 chances)\n3. Hard (3 chances)")

		fmt.Println("")

		fmt.Println("Enter your choice:")

		fmt.Scanln(&level)

		fmt.Printf("Great! You have selected the %s difficulty level.\n", m[level])
		fmt.Println("Let's start the game!")

		for {
			count++
			fmt.Println("Enter your guess:")
			fmt.Scanln(&tempNums)
			if tempNums == nums {
				fmt.Printf("Congratulations! You guessed the correct number in %d attempts.\n", count)
				break
			} else if nums > tempNums {
				fmt.Printf("Incorrect! The number is greater than %d.\n", tempNums)
			} else {
				fmt.Printf("Incorrect! The number is less than %d.\n", tempNums)
			}

			if (level == 1 && count >= 10) || (level == 2 && count >= 5) || (level == 3 && count >= 3) {
				fmt.Printf("机会用完，本局游戏结束,答案：%d\n", nums)
				break
			}
		}
		count = 0
		fmt.Println("是否继续? 1:开启下一局，0：结束游戏")
		var isContinue bool
		fmt.Scanln(&isContinue)
		if !isContinue {
			fmt.Println("Game Over\nBye Bye!!!")
			break
		}
	}

}
