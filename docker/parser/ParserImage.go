package parser

import (
	"encoding/json"
	"strings"

	"jdss.docker.client/docker/dto"
)

/*
ContainerDto used to parse result
*/
type ParserImageCommand struct {
}

/*
	Function to parse a string (Json format) to ImageDto
*/
func (parser *ParserImageCommand) Parse(result string) []dto.ImageDto {
	var images []dto.ImageDto
	json.Unmarshal([]byte(result), &images)

	for i := 0; i < len(images); i++ {
		images[i].ID = strings.TrimPrefix(images[i].ID, "sha256:")
	}
	return images

}
