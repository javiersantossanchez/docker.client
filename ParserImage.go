package main

import "encoding/json"

/*
ContainerDto used to parse result
*/
type ParserImageCommand struct {
}

func (parser *ParserImageCommand) Parse(result string) []ImageDto {
	var bird []ImageDto
	json.Unmarshal([]byte(result), &bird)
	return bird

}
