package main

import (
	"fmt"
	"os"

	"github.com/joshjon/toyrobot/internal/robot"
)

func main() {
	fmt.Println("Enter commands:")
	errs := robot.RunSimulation(os.Stdin, os.Stdout)
	for err := range errs {
		fmt.Println(err)
	}
}
