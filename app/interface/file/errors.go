package file

type fileInputError string

func (e fileInputError) Error() string { return string(e) }

const (
	ErrReadingToFile                = fileInputError("error reading to file")
	ErrInstructionNotCompleted      = fileInputError("error Sorry but the instructions should have to instructions for the plane and a robot")
	ErrAnInstructionIsMissing       = fileInputError("error sorry an instruction is  missing")
	ErrMappingStringToSurfaceEntity = fileInputError("error mapping string to surface entity")
)
