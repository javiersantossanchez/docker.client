package parser

import (
	"encoding/json"

	"jdss.docker.client/docker/dto"
)

/*
ContainerDto used to parse result
*/
type ParserContainerCommand struct {
}

func (parser *ParserContainerCommand) Parse(result string) []dto.ContainerDto {
	var bird []dto.ContainerDto
	json.Unmarshal([]byte(result), &bird)
	return bird

}
