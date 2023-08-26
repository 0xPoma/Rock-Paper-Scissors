package main

import "fmt"
import "time"
import "math/rand"

func main(){
	var choice int
	fmt.Println("--Rock,Paper,Scissors--")
	fmt.Println(">>1.Rock 2.Paper 3.Scissors")
	fmt.Scan(&choice)
	rand.Seed(time.Now().UnixNano())
	intchoice:=rand.Intn(3)
	switch choice {
		case 1: // Rock
			if (intchoice == 0){
				fmt.Println("Draw")
			}
			if(intchoice == 1){ //Paper
				fmt.Println("You lose")
			}
			if (intchoice == 2){ 
				fmt.Println("You win") //Scissors
			}
		case 2: //Paper
			if (intchoice == 0){
				fmt.Println("You win")
			}
			if(intchoice == 1){
				fmt.Println("Draw")
			}
			if (intchoice == 2){
				fmt.Println("You lose")
			}
		case 3: //Scissors
			if (intchoice == 0) { 
				fmt.Println("You lose")
			}
			if (intchoice == 1) {
				fmt.Println("You win")
			}
			if (intchoice == 2){
				fmt.Println("Draw")
			}
		default:
			fmt.Println("Something went wrong")
			}
}