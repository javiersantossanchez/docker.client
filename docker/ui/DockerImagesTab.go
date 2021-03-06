package ui

import (
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"jdss.docker.client/docker"
	"jdss.docker.client/docker/com"
	"jdss.docker.client/docker/parser"
)

// GetImageTab Hello test
func GetImageTab(viewDetailCallback func(id string)) *widget.TabItem {
	imageCommand := com.ListImagesCommand{Docker: docker.DockerConnector{}}
	imageResult := imageCommand.Execute()
	parserImages := parser.ParserImagesCommand{}
	images := parserImages.Parse(imageResult)

	containerTag := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), widget.NewLabel("Image Tag"), layout.NewSpacer())

	containerID := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), widget.NewLabel("Image ID"), layout.NewSpacer())

	containerSize := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), widget.NewLabel("Image Size"), layout.NewSpacer())

	containerDetail := fyne.NewContainerWithLayout(layout.NewVBoxLayout(), widget.NewLabel(""), layout.NewSpacer())

	for _, container := range images {
		if len(container.RepoTags) >= 1 {
			containerTag.AddObject(widget.NewLabel(container.RepoTags[0]))
		} else {
			containerTag.AddObject(widget.NewLabel("-"))
		}

		containerID.AddObject(widget.NewLabel(container.ID))
		containerSize.AddObject(widget.NewLabel(strconv.Itoa(container.Size)))

		idValue := container.ID

		onClick := func() {
			viewDetailCallback(idValue)
		}

		viewDetailButton := widget.NewButton("View Detail", onClick)
		containerDetail.AddObject(viewDetailButton)

	}

	c := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), containerTag, layout.NewSpacer(),
		containerID, layout.NewSpacer(), containerSize, layout.NewSpacer(), containerDetail)

	dockerImages := widget.NewVScrollContainer(c)
	dockerImages.SetMinSize(fyne.Size{Height: 620, Width: 580})
	imageListTab := widget.NewTabItem("Images List", dockerImages)
	return imageListTab
}
