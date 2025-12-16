package domain

import "fmt"

type Command interface {
	Execute(robot *Robot, world *World)
}

type LeftCommand struct{}

func (c LeftCommand) Execute(robot *Robot, world *World) {
	fmt.Println("left")
	robot.Orientation = robot.Orientation.Left()
}

type RightCommand struct{}

func (c RightCommand) Execute(robot *Robot, world *World) {
	fmt.Println("right")
	robot.Orientation = robot.Orientation.Right()
}

type ForwardCommand struct{}

func (c ForwardCommand) Execute(robot *Robot, world *World) {
	fmt.Println("forward")
}
