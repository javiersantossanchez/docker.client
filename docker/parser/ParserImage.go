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
func (parser *ParserImageCommand) Parse(result string) dto.ImageDetailDto {
	var image dto.ImageDetailDto
	json.Unmarshal([]byte(result), &image)

	image.ID = strings.TrimPrefix(image.ID, "sha256:")
	image.Layers = image.RootFS.Layers
	return image

}
