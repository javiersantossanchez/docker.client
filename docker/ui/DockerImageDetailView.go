package ui

import (
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"jdss.docker.client/docker/dto"
)

/**
Hello test
**/
func GetImageDetailView(image dto.ImageDetailDto, containers []dto.ContainerDto) *fyne.Container {
	container := fyne.NewContainerWithLayout(layout.NewFormLayout(),
		widget.NewLabel("Image ID"), widget.NewLabel(image.ID),
		widget.NewLabel("Image size"), widget.NewLabel(strconv.Itoa(image.Size)),
		widget.NewLabel("Author"), widget.NewLabel(image.Author),
	)

	container.AddObject(widget.NewLabel("Container"))
	container.AddObject(widget.NewLabel(""))
	for _, curent := range containers {

		container.AddObject(widget.NewLabel(curent.Names[0]))
		container.AddObject(widget.NewLabel(curent.State))
	}

	return container

}
