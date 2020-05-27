package main

import (
	"log"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"jdss.docker.client/docker"
	"jdss.docker.client/docker/com"
	"jdss.docker.client/docker/parser"
)

func main() {

	imageCommand := com.ImageCommand{Docker: docker.DockerConnector{}}
	imageResult := imageCommand.Execute()
	parserImage := parser.ParserImageCommand{}
	images := parserImage.Parse(imageResult)

	containerCommand := com.ListContainerCommand{Docker: docker.DockerConnector{}}
	containerResult := containerCommand.Execute()
	parse := parser.ParserContainerCommand{}
	bird := parse.Parse(containerResult)

	app := app.New()
	w := app.NewWindow("Hello")

	rt := widget.NewVBox()

	for _, container := range images {
		imageIDLabel := widget.NewLabel(container.ID)
		imageSizelabel := widget.NewLabel(strconv.Itoa(container.Size))
		idValue := container.ID
		onClick := func() {

			container := widget.NewVBox()
			container.Append(widget.NewLabel(idValue))
			//widget

			as := widget.NewModalPopUp(container, w.Canvas())
			container.Append(widget.NewButton("asas", func() {
				log.Println("tapped" + idValue)
				as.Hide()
			}))

		}

		viewDetailButton := widget.NewButton("View Detail", onClick)

		rt.Append(widget.NewHBox(imageIDLabel, imageSizelabel, viewDetailButton))
	}

	dockerImages := widget.NewVScrollContainer(
		rt,
	)
	dockerImages.SetMinSize(fyne.Size{Height: 120, Width: 580})

	imageListTab := widget.NewTabItem("Images List", dockerImages)

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

	containerListTab := widget.NewTabItem("Container List", dockerContainers)

	tabs := widget.NewTabContainer(
		containerListTab,
		imageListTab,
	)

	w.SetContent(widget.NewVBox(
		tabs,
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))
	w.Resize(fyne.Size{Height: 520, Width: 680})

	w.ShowAndRun()
}
