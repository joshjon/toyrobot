package main

import (
	"fmt"
	"os"

	"github.com/joshjon/toyrobot/internal/simulation"
)

func main() {
	fmt.Println("Enter commands:")
	reader := simulation.NewReader(os.Stdin)
	executor := simulation.NewExecutor()
	simulator := simulation.NewSimulator(reader, executor, os.Stdout)
	errs := simulator.Run()
	for err := range errs {
		if err != nil {
			fmt.Println(err)
		}
	}
}
