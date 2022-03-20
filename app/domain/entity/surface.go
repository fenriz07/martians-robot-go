package entity

type Surface struct {
	Coordinates
	LostCoordinates []Coordinates
}

func (s *Surface) RobotIsLost(robot Robot) bool {

	if robot.Coordinates.X > s.X {
		s.LostCoordinates = append(s.LostCoordinates, robot.Coordinates)
		return true
	}

	if robot.Coordinates.Y > s.Y {
		s.LostCoordinates = append(s.LostCoordinates, robot.Coordinates)
		return true
	}

	if robot.Coordinates.X < 0 {
		s.LostCoordinates = append(s.LostCoordinates, robot.Coordinates)
		return true
	}

	if robot.Coordinates.Y < 0 {
		s.LostCoordinates = append(s.LostCoordinates, robot.Coordinates)
		return true
	}

	return false
}
