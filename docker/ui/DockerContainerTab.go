package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"jdss.docker.client/docker"
	"jdss.docker.client/docker/com"
	"jdss.docker.client/docker/parser"
)

func GetContainerTab() *widget.TabItem {
	containerCommand := com.ListContainerCommand{Docker: docker.DockerConnector{}}
	containerResult := containerCommand.Execute()
	parse := parser.ParserContainerCommand{}
	containers := parse.Parse(containerResult)

	rt := widget.NewVBox()
	for _, container := range containers {

		containerIDLabel := widget.NewLabel(container.ID)
		imagelabel := widget.NewLabel(container.Image)

		rt.Append(widget.NewHBox(containerIDLabel, imagelabel))
	}
	dockerContainers := widget.NewVScrollContainer(
		rt,
	)

	dockerContainers.SetMinSize(fyne.Size{Height: 120, Width: 580})

	containerListTab := widget.NewTabItem("Images List", dockerContainers)

	return containerListTab

}
