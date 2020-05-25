package main

import (
	"io/ioutil"
	"log"
)

/*
ContainerDto used to parse result
*/
type ListContainerCommand struct {
	docker DockerConnector
}

func (command *ListContainerCommand) execute() string {
	httpConnector := command.docker.GetConnector()

	resp, err := httpConnector.Get(command.docker.baseUrl() + "/containers/json")
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
