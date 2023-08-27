package main

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"math/rand"
	"os"
	"time"
)

func main() {
	var choice, randchoice, n, menu int
	var countDraw, countWin, countLose int
	fmt.Println("--Rock,Paper,Scissors--")
	// fmt.Println(">>1.Rock 2.Paper 3.Scissors")
	fmt.Println("1.Play the game with bot")
	fmt.Println("2.Check the result of game")
	fmt.Scan(&menu)
	_ = os.Remove("bar.html")
	rand.Seed(time.Now().UnixNano())
	switch menu {
	case 1:
		fmt.Println(">>1.Rock 2.Paper 3.Scissors")
		fmt.Scan(&choice)
		intchoice := rand.Intn(3)
		switch choice {
		case 1: // Rock
			if intchoice == 0 {
				fmt.Println("Draw")
			}
			if intchoice == 1 { //Paper
				fmt.Println("You lose")
			}
			if intchoice == 2 {
				fmt.Println("You win") //Scissors
			}
		case 2: //Paper
			if intchoice == 0 {
				fmt.Println("You win")
			}
			if intchoice == 1 {
				fmt.Println("Draw")
			}
			if intchoice == 2 {
				fmt.Println("You lose")
			}
		case 3: //Scissors
			if intchoice == 0 {
				fmt.Println("You lose")
			}
			if intchoice == 1 {
				fmt.Println("You win")
			}
			if intchoice == 2 {
				fmt.Println("Draw")
			}
		default:
			fmt.Println("Something went wrong")
		}

	case 2:
		fmt.Println("Enter the number of games you wanna to see")
		fmt.Scan(&n)
		defer fmt.Println("The file was created")
		for i := 0; i < n; i++ {
			randchoice = rand.Intn(3) + 1
			intchoice := rand.Intn(3)
			switch randchoice {
			case 1: // Rock
				if intchoice == 0 {
					countDraw++
				}
				if intchoice == 1 { //Paper
					countLose++
				}
				if intchoice == 2 {
					//Scissors
					countWin++
				}
			case 2: //Paper
				if intchoice == 0 {
					countWin++
				}
				if intchoice == 1 {
					countDraw++
				}
				if intchoice == 2 {
					countLose++
				}
			case 3: //Scissors
				if intchoice == 0 {
					countLose++
				}
				if intchoice == 1 {
					countWin++
				}
				if intchoice == 2 {
					countDraw++
				}
			default:
				fmt.Println("Something went wrong")
			}
		}
		bar := charts.NewBar()
		bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
			Title: "Comparing ",
		}))

		bar.SetXAxis([]string{"Win", "Lose", "Draw"}).
			AddSeries("Numbers", getDiagramm(countWin, countLose, countDraw))
		f, _ := os.Create("bar.html")
		bar.Render(f)
	}
}
func getDiagramm(countWin, countLose, countDraw int) []opts.BarData {
	values := []opts.BarData{
		{Value: countWin},
		{Value: countLose},
		{Value: countDraw},
	}
	return values
}
