package com

import (
	"io/ioutil"
	"log"

	"jdss.docker.client/docker"
)

/*
ContainerDto used to parse result
*/
type ImageCommand struct {
	Docker docker.DockerConnector
}

/*
sdadasdads
*/
func (command *ImageCommand) Execute() string {
	httpConnector := command.Docker.GetConnector()

	resp, err := httpConnector.Get(command.Docker.BaseUrl() + "/images/json")
	if err != nil {
		log.Fatal(err)
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)

	resp.Body.Close()

	return responseString

}
