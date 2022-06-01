package main

import (
	"cezhang/rangelist/rangelist"
)

func main() {

	rl := rangelist.New()
	rl.Add([2]int{1, 5})
	rl.Print()
	rl.Add([2]int{10, 20})
	rl.Print()
	rl.Add([2]int{20, 20})
	rl.Print()
	rl.Add([2]int{20, 21})
	rl.Print()
	rl.Add([2]int{2, 4})
	rl.Print()
	rl.Add([2]int{3, 8})
	rl.Print()
	rl.Remove([2]int{10, 10})
	rl.Print()

	rl.Remove([2]int{10, 11})
	rl.Print()
	rl.Remove([2]int{15, 17})
	rl.Print()
	rl.Remove([2]int{3, 19})
	rl.Print()
}
