package main

import (
	"fmt"
	"image/color"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Window testing")

	w.Resize(fyne.NewSize(2000, 1000))
	napis := widget.NewLabel("")

	przycisk := widget.NewButton("Button Testing", func() {
		fmt.Println("Button Output")
	})
	hak := widget.NewCheck("True", func(b bool) {
		if b {
			fmt.Println("Male")
		} else {
			fmt.Println("female")
		}
	})

	url, _ := url.Parse("https://github.com")
	link := widget.NewHyperlink("visit me:", url)

	colorx := color.NRGBA{R: 100, G: 20, B: 150, A: 255}
	testx := canvas.NewText("Color text", colorx)

	// Use a container to hold all widgets
	content := container.NewVBox(napis, przycisk, hak, link, testx)
	w.SetContent(content)

	w.ShowAndRun()
}
