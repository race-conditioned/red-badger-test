package simulator

import "red-badger-test/internal/domain"

type Simulator struct{}

func (s *Simulator) RunRobot(robot *domain.Robot, world *domain.World, commands []domain.Command) {
	for _, cmd := range commands {
		cmd.Execute(robot, world)
		if robot.Lost {
			break
		}
	}
}
