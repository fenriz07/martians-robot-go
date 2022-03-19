package main

import (
	"martians/app/interface/file"
	"martians/app/usecase"
)

func main() {

	nameFile := "input"

	useCaseExecuteSequence := usecase.ExecuteSequenceUseCase{}

	input := file.NewInputInterface(useCaseExecuteSequence.Launch)

	input.StartInstructions(nameFile)
}
