package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"jdss.docker.client/docker"
	"jdss.docker.client/docker/com"
	"jdss.docker.client/docker/parser"
	"jdss.docker.client/docker/ui"
)

func main() {

	app := app.New()
	w := app.NewWindow("Hello")

	tab := widget.NewTabContainer(
		ui.GetImageTab(func(id string) {

			imageCommand := com.ImageCommand{Docker: docker.DockerConnector{}, Context: map[string]string{"imageID": id}}
			imageResult := imageCommand.Execute()
			parserImage := parser.ParserImageCommand{}
			image := parserImage.Parse(imageResult)

			xxx := com.ListContainerByImageCommand{Docker: docker.DockerConnector{}, Context: map[string]string{"imageID": id}}
			value := xxx.Execute()
			par := parser.ParserContainerCommand{}
			result := par.Parse(value)

			container := widget.NewVBox()
			container.Append(ui.GetImageDetailView(image, result))

			as := widget.NewModalPopUp(container, w.Canvas())
			container.Append(widget.NewButton("asas", func() {
				as.Hide()
			}))
		}),
		ui.GetContainerTab(),
	)

	// content := widget.NewVBox(
	// 	//tabs,
	// 	widget.NewButton("Quit", func() {
	// 		app.Quit()
	// 	}),
	// )

	//tab.Resize(fyne.Size{Height: 1520, Width: 400})
	container := widget.NewVBox(tab)

	w.SetContent(container)
	w.CenterOnScreen()

	w.ShowAndRun()
}
