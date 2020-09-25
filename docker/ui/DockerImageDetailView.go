package ui

import (
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"jdss.docker.client/docker/dto"
)

// GetImageDetailView d
func GetImageDetailView(image dto.ImageDetailDto, containers []dto.ContainerDto) *fyne.Container {

	mainContainer := fyne.NewContainerWithLayout(layout.NewVBoxLayout())

	formContainer := fyne.NewContainerWithLayout(layout.NewFormLayout(),
		widget.NewLabel("Image Tag"), widget.NewLabel(image.RepoTags[0]),
		widget.NewLabel("Image ID"), widget.NewLabel(image.ID),
		widget.NewLabel("Image size"), widget.NewLabel(strconv.Itoa(image.Size)),
		widget.NewLabel("Author"), widget.NewLabel(image.Author),
	)
	mainContainer.AddObject(formContainer)
	mainContainer.AddObject(widget.NewLabel("Container"))

	containerName := fyne.NewContainerWithLayout(layout.NewVBoxLayout())
	containerStatus := fyne.NewContainerWithLayout(layout.NewVBoxLayout())

	if containers == nil || len(containers) > 0 {
		for _, curent := range containers {
			containerName.AddObject(widget.NewLabel(curent.Names[0]))
			containerStatus.AddObject(widget.NewLabel(curent.State))

		}

		cont := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), containerName, containerStatus)
		scrollContainer := widget.NewVScrollContainer(cont)
		scrollContainer.SetMinSize(fyne.Size{Height: 120, Width: 580})
		mainContainer.AddObject(scrollContainer)
	}

	mainContainer.AddObject(widget.NewLabel("Layers"))
	containerLayers := fyne.NewContainerWithLayout(layout.NewVBoxLayout())
	if image.Layers == nil || len(image.Layers) > 0 {
		for _, layer := range image.Layers {
			containerLayers.AddObject(widget.NewLabel(layer))
		}
		cont := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), containerLayers)
		scrollContainer := widget.NewVScrollContainer(cont)
		scrollContainer.SetMinSize(fyne.Size{Height: 220, Width: 580})
		mainContainer.AddObject(scrollContainer)
	}

	return mainContainer

}
