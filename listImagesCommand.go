package main

import (
	"io/ioutil"
	"log"
)

/*
ContainerDto used to parse result
*/
type ListImagesCommand struct {
	docker DockerConnector
}

func (command *ListImagesCommand) execute() string {
	httpConnector := command.docker.GetConnector()

	resp, err := httpConnector.Get(command.docker.baseUrl() + "/images/json")
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
