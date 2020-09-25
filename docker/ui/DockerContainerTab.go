package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"jdss.docker.client/docker"
	"jdss.docker.client/docker/com"
	"jdss.docker.client/docker/parser"
)

//GetContainerTab test
func GetContainerTab() *widget.TabItem {
	containerCommand := com.ListContainerCommand{Docker: docker.DockerConnector{}}
	containerResult := containerCommand.Execute()
	parse := parser.ParserContainerCommand{}
	containers := parse.Parse(containerResult)

	containerID := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), widget.NewLabel("Container ID"), layout.NewSpacer())
	containerName := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), widget.NewLabel("Container Name"), layout.NewSpacer())

	containerStatus := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), widget.NewLabel("Container Status"), layout.NewSpacer())

	containerImage := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), widget.NewLabel("Container Image"), layout.NewSpacer())

	for _, container := range containers {

		containerID.AddObject(widget.NewLabel(container.ID))
		containerName.AddObject(widget.NewLabel(container.Names[0]))
		containerImage.AddObject(widget.NewLabel(container.Image))
		containerStatus.AddObject(widget.NewLabel(container.Status))

	}
	c := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), containerID, layout.NewSpacer(), containerName, layout.NewSpacer(), containerStatus, layout.NewSpacer(), containerImage)

	dockerImages := widget.NewVScrollContainer(c)
	dockerImages.SetMinSize(fyne.Size{Height: 620, Width: 580})
	imageListTab := widget.NewTabItem("Container List", dockerImages)
	return imageListTab

}
