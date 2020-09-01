package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRand() int {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100)
	return num
}

func judge(t int, i *int, limit int) bool {
	if t == *i {
		win()
		return true
	} else if t > *i {
		fmt.Printf("%v is smaller than the target😢\n", *i)
	} else {
		fmt.Printf("%v is larger than the target😢\n", *i)
	}

	fmt.Printf("You have %v times left. Please try again!\n", limit)
	limit--
	if limit < 0 {
		lose()
		return false
	}
	_, _ = fmt.Scanln(i)
	return judge(t, i, limit)
}

func win() {
	fmt.Println("Congratulations! You are right!👏")
}

func lose() {
	fmt.Println("Sorry, you have use all your chance☠")
	fmt.Println("Would you like to play again ?(Y or N)")
	var label string
	_, _ = fmt.Scanln(&label)
	if label == "Y" {
		start()
	} else {
		fmt.Println("---------- Bye! ----------")
	}
}

func start() {
	fmt.Println("---------- Guess Number Game Start ----------")
	fmt.Println("You have seven times to try, please input a number!")
	// 1. generate a random Inter
	target := generateRand()
	// 2. define a current to store the user input by transfer thr pointer
	var current int
	_, _ = fmt.Scanln(&current)
	// 3. we limit the user-input-count
	limit := 7
	// judge the target and current
	judge(target, &current, limit)
}

func main() {
	start()
}
