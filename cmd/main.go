package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/joshjon/toyrobot/internal/robot"
)

func main() {
	fmt.Println("Enter commands:")
	if err := run(os.Stdin); err != nil {
		fmt.Printf("unexpected error occurred: %v\n", err)
		os.Exit(1)
	}
}

func run(reader io.Reader) error {
	commands := make(chan robot.Command, 10)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		robot.RunSimulation(commands)
		wg.Done()
	}()

	if err := readCommands(reader, commands); err != nil {
		return err
	}

	wg.Wait()
	return nil
}

func readCommands(reader io.Reader, commands chan robot.Command) error {
	defer close(commands)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return err
		}
		line := scanner.Text()
		command, err := robot.CommandFromString(line)
		if err != nil {
			// Print command validation errors to inform the user of bad input
			fmt.Println(err.Error())
			continue
		}
		commands <- command
	}

	return nil
}
