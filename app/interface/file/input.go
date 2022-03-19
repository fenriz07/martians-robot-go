package file

import (
	"bufio"
	"martians/app/interface/file/mappers"
	"martians/app/share/constants"
	"martians/app/usecase"
	"os"

	"github.com/rs/zerolog/log"
)

const (
	StartInstructionsNameFunction = "StartInstructions"
)

type input struct {
	usecase.ExecuteSequence
}

func NewInputInterface(u usecase.ExecuteSequence) input {
	return input{
		ExecuteSequence: u,
	}
}

func (i *input) StartInstructions(nameFile string) (interface{}, error) {
	lines, err := readFile(nameFile)

	if err != nil {
		log.Error().Err(err).
			Str(constants.Function, StartInstructionsNameFunction).
			Msg(ErrReadingToFile.Error())

		return nil, ErrReadingToFile

	}
	lenLines := len(lines)

	if lenLines < 3 {
		log.Error().
			Str(constants.Function, StartInstructionsNameFunction).
			Msg(ErrInstructionNotCompleted.Error())

		return nil, ErrInstructionNotCompleted
	}

	if lenLines%2 == 0 {
		log.Error().
			Str(constants.Function, StartInstructionsNameFunction).
			Msg(ErrAnInstructionIsMissing.Error())

		return nil, ErrAnInstructionIsMissing
	}

	surface, errMappingStringToSurface := mappers.StringToSurfaceEntity(lines[0])

	if errMappingStringToSurface != nil {
		return nil, ErrMappingStringToSurfaceEntity
	}

	allInstructionsWithOutSurface := lines[1:]

	allRobotsAndItsInstruction := chunkSlice(allInstructionsWithOutSurface, 2)

	robots, err := mappers.MapStringToRobotEntity(allRobotsAndItsInstruction)

	if err != nil {

		log.Error().
			Err(err).
			Str(constants.Function, StartInstructionsNameFunction).
			Msg("Error mapping map robotWithInstruction to robot entity")

	}

	_, _ = i.ExecuteSequence(robots, *surface)

	return nil, nil
}

func readFile(nameFile string) ([]string, error) {
	file, err := os.Open(nameFile)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func chunkSlice(slice []string, chunkSize int) [][]string {
	var chunks [][]string
	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}
