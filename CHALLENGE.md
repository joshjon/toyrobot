# Toy Robot Simulation Coding Challenge

## Overview

This challenge involves creating a simulation of a toy robot moving on a square tabletop of dimensions 5x5 units. Your task is to develop an application that can read and execute specific commands, maintaining the toy robot within the confines of the tabletop.

## Requirements

- The robot should move within a 5x5 square tabletop.
- No obstructions exist on the tabletop surface.
- Prevent the robot from "falling" off the tabletop.
- The robot should interpret the following commands:
    - PLACE X,Y,F
    - MOVE
    - LEFT
    - RIGHT
    - REPORT

**Command Details:**

- PLACE will put the robot on the table at position X,Y and facing NORTH, SOUTH, EAST or WEST.
- The origin (0,0) is the SOUTH WEST most corner.
- The first valid command is a PLACE command.
- Commands can be issued in any order, including another PLACE command.
- MOVE will move the robot one unit forward in the direction it is currently facing.
- LEFT and RIGHT will rotate the robot 90 degrees in the specified direction without changing its position.
- REPORT will announce the X,Y and F of the robot through the standard output.

## Constraints

Prevent the robot from falling off the table. This also includes the initial placement of the toy robot. Any movement that would cause the robot to fall must be ignored.

## Sample Input and Output

### a) Scenario 1

```
PLACE 0,0,NORTH
MOVE
REPORT
```

Output: `0,1,NORTH`

### b) Scenario 2

```
PLACE 0,0,NORTH
LEFT
REPORT
```

Output: `0,0,WEST`

## Deliverables

- GitHub repository containing your solution in your preferred programming language.
- Instructions on how to run your project.
- Unit tests demonstrating the functionality of your application.

This can be a Command Line Interface (CLI) application OR a User Interface (UI) application that showcases your skills. The emphasis is on the implementation of the logic with clean, readable code.

## **What we would like to see**

1. **Functionality:** Your solution should correctly implement the toy robot simulation as per the given requirements and constraints.
2. **Code Quality:** Your code should be clean, well-structured, and easy to read. Make good use of comments where necessary, but we’d like to code that’s self-describing.
3. **Testing:** Include as many tests as you deem appropriate to demonstrate that your solution works correctly. We don’t follow strict TDD at Mr Yum, and we’re interested in hearing your opinions on testing.
4. **Problem Solving:** We're interested in how you approach problems and how you solve them. There are no right answers here - only tradeoffs. So, feel free to document them.

## The most successful candidates usually

1. Write a good README.
2. Make the tradeoffs clear in the code, or in the README.
3. Have consistency in naming, and structuring their code.
4. Don’t over-engineer their solution.
    1. Instead, they can tell what they could’ve done, if afforded more time.
5. Code in the language they’re most comfortable in.
