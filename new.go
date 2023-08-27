package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func main() {
	var choice int
	var countDraw, countWin, countLose int
	fmt.Println("--Rock,Paper,Scissors--")
	fmt.Println(">>1.Rock 2.Paper 3.Scissors")
	// fmt.Scan(&choice)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 1000; i++ {
		choice = rand.Intn(3) + 1
		intchoice := rand.Intn(3)
		switch choice {
		case 1: // Rock
			if intchoice == 0 {
				// fmt.Println("Draw")
				countDraw++
			}
			if intchoice == 1 { //Paper
				// fmt.Println("You lose")
				countLose++
			}
			if intchoice == 2 {
				// fmt.Println("You win") //Scissors
				countWin++
			}
		case 2: //Paper
			if intchoice == 0 {
				// fmt.Println("You win")
				countWin++
			}
			if intchoice == 1 {
				// fmt.Println("Draw")
				countDraw++
			}
			if intchoice == 2 {
				// fmt.Println("You lose")
				countLose++
			}
		case 3: //Scissors
			if intchoice == 0 {
				// fmt.Println("You lose")
				countLose++
			}
			if intchoice == 1 {
				// fmt.Println("You win")
				countWin++
			}
			if intchoice == 2 {
				// fmt.Println("Draw")
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

func getDiagramm(countWin, countLose, countDraw int) []opts.BarData {
	values := []opts.BarData{
		{Value: countWin},
		{Value: countLose},
		{Value: countDraw},
	}
	return values
}
