package com

import (
	"io/ioutil"
	"log"

	"jdss.docker.client/docker"
)

/*
ContainerDto used to parse result
*/
type ListContainerByImageCommand struct {
	Docker  docker.DockerConnector
	Context map[string]string
}

func (command *ListContainerByImageCommand) Execute() string {
	httpConnector := command.Docker.GetConnector()
	resp, err := httpConnector.Get(command.Docker.BaseUrl() + "/containers/json?filters={\"ancestor\":[\"" + command.Context["imageID"] + "\"]}")
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
