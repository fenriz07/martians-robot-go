package usecase

import "martians/app/domain/entity"

type ExecuteSequence func(robots []entity.Robot, surface entity.Surface) (interface{}, error)

type ExecuteSequenceUseCase struct {
}

func (u *ExecuteSequenceUseCase) Launch(r []entity.Robot, s entity.Surface) (interface{}, error) {

	robots := r
	surface := s

	for k, robot := range robots {

		if isLost := surface.RobotIsLost(robot); isLost {
			robots[k].Lost = isLost

			continue
		}

		/*for _, i := range robot.Instructions.Instructions {

		}*/
	}

	return nil, nil
}
