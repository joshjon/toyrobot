# 🤖 Toy Robot

Toy robot [coding challenge](CHALLENGE.md) solution written in Go.

## 🧰 Dependencies

- [Go 1.20](https://go.dev/dl/)

## 🚀 Run Simulation

1. Start the toy robot simulator
   <br><br>
  
   Go:

    ```shell
    go run cmd/main/main.go
    ```
   Docker:

     ```shell
     docker build -t local/toyrobot .
     docker run --interactive --rm --name toyrobot local/toyrobot
     ```

2. Enter commands e.g.

    ````
    PLACE 0,0,NORTH
    MOVE
    REPORT
    ````

3. Stop the toy robot simulator with `ctrl+D`

## 🔬 Run Tests

```shell
go test ./... -cover
```

## 📝 Domain Model

```mermaid
classDiagram
    direction LR

   class Reader {
      -reader io.Reader
      +Run(commands chan<- Command) <-chan error
   }

   class Executor {
      -reader io.Reader
      +Run(commands <-chan Command) <-chan error
   }

   class Simulator {
      -reader Reader
      -executor Executor
      +Run() <-chan error
   }

    class Command {
        <<interface>>
        +Execute(state *State)
    }

    class CommandPlace {
        +X int
        +Y int
        +Direction Direction
        +Execute(state *State)
    }

    class CommandMove {
        +Execute(state *State)
    }

    class CommandRotateLeft {
        +Execute(state *State)
    }

    class CommandRotateRight {
        +Execute(state *State)
    }

    class CommandReport {
        +Execute(state *State)
    }

    class State {
        -maxX int
        -maxY int
        -posX int
        -posY int
        -direction Direction
        -placed bool
        -writer io.Writer
        -report() error
    }

    class Direction {
        <<enumeration>>
        +North
        +South
        +East
        +West
        +Left()
        +Right()
        +Axes() int,int
        +String() string
    }


    Simulator -- "1" Reader: has
    Simulator -- "1" Executor: has
    Simulator -- "1" State: has
    Executor --> Command: executes
    Command <|.. CommandPlace: implements
    Command <|.. CommandMove: implements
    Command <|.. CommandRotateLeft: implements
    Command <|.. CommandRotateRight: implements
    Command <|.. CommandReport: implements
```
