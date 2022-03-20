package usecase

import "martians/app/domain/entity"

type ExecuteSequence func(robots []entity.Robot, surface entity.Surface) []entity.Robot

type ExecuteSequenceUseCase struct {
}

func (u *ExecuteSequenceUseCase) Launch(r []entity.Robot, s entity.Surface) []entity.Robot {

	robots := r
	surface := s

	for k, robot := range robots {

		if isLost := surface.RobotIsLost(robot); isLost {
			robots[k].Lost = isLost

			continue
		}

		for _, i := range robot.Instructions.Instructions {

			if entity.L == i {
				robots[k].Orientation = changeOrientationLeft(robots[k].Orientation)

				continue
			}

			if entity.R == i {
				robots[k].Orientation = changeOrientationRight(robots[k].Orientation)
				continue
			}

			if entity.F == i {
				robots[k].Coordinates = foward(robots[k], surface)

				if isLost := surface.RobotIsLost(robots[k]); isLost {
					robots[k].Lost = isLost
				}

				continue
			}

		}
	}

	return robots
}

func foward(r entity.Robot, s entity.Surface) entity.Coordinates {

	robot := r

	if robot.Orientation == entity.N {
		robot.Coordinates.Y = robot.Coordinates.Y + 1
	}

	if robot.Orientation == entity.S {
		robot.Coordinates.Y = robot.Coordinates.Y - 1
	}

	if robot.Orientation == entity.E {
		robot.Coordinates.X = robot.Coordinates.X + 1
	}

	if robot.Orientation == entity.W {
		robot.Coordinates.X = robot.Coordinates.X - 1
	}

	for _, c := range s.LostCoordinates {
		if c.X == robot.Coordinates.X && c.Y == robot.Coordinates.Y {
			return r.Coordinates
		}
	}

	return robot.Coordinates
}

func changeOrientationLeft(o entity.Orientation) entity.Orientation {
	n := o.EnumIndex()

	n = n - 1

	if n == 0 {
		n = 4
	}

	return entity.NewOrientationtByInt(n)
}

func changeOrientationRight(o entity.Orientation) entity.Orientation {
	n := o.EnumIndex()

	n = n + 1

	if n == 5 {
		n = 1
	}

	return entity.NewOrientationtByInt(n)
}
