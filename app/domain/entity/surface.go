package entity

type Surface struct {
	Coordinates
}

func (s *Surface) RobotIsLost(robot Robot) bool {

	if robot.Coordinates.X > s.X {
		return true
	}

	if robot.Coordinates.Y > s.Y {
		return true
	}

	if robot.Coordinates.X < 0 {
		return true
	}

	if robot.Coordinates.Y < 0 {
		return true
	}

	return false
}
