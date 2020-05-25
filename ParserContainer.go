package main

import "encoding/json"

/*
ContainerDto used to parse result
*/
type ParserContainerCommand struct {
}

func (parser *ParserContainerCommand) Parse(result string) []ContainerDto {
	var bird []ContainerDto
	json.Unmarshal([]byte(result), &bird)
	return bird

}
