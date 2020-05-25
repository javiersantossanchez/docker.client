package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"jdss.docker.client/docker"
	"jdss.docker.client/docker/com"
	"jdss.docker.client/docker/parser"
)

func main() {

	imageCommand := com.ListImagesCommand{Docker: docker.DockerConnector{}}
	imageResult := imageCommand.Execute()
	parserImage := parser.ParserImageCommand{}
	images := parserImage.Parse(imageResult)

	containerCommand := com.ListContainerCommand{Docker: docker.DockerConnector{}}
	containerResult := containerCommand.Execute()
	parse := parser.ParserContainerCommand{}
	bird := parse.Parse(containerResult)

	app := app.New()

	sizeImage := len(images)
	var dddImages = make([]fyne.CanvasObject, sizeImage)

	indexImages := 0
	for _, container := range images {
		dddImages[indexImages] = widget.NewLabel(container.ID + "......")
		indexImages++
	}

	rt := widget.NewGroup(
		"Images",
		dddImages...,
	)

	dockerImages := widget.NewVScrollContainer(
		rt,
	)
	dockerImages.SetMinSize(fyne.Size{Height: 120, Width: 580})

	size := len(bird)

	var ddd = make([]fyne.CanvasObject, size)

	index := 0

	for _, container := range bird {
		ddd[index] = widget.NewLabel(container.ID + "......" + container.Image)
		index++
	}

	dockerContainers := widget.NewGroup("Containers",
		ddd...,
	)

	w := app.NewWindow("Hello")

	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		dockerImages,
		dockerContainers,
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))
	w.Resize(fyne.Size{Height: 520, Width: 680})

	w.ShowAndRun()
}
