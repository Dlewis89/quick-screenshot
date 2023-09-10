package main

import (
	"fmt"
	"image/png"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/kbinani/screenshot"
)

func takeScreenshot(displays []Display, selectedScreen string) {
	for i := 0; i < len(displays); i++ {
		if displays[i].name != selectedScreen {
			continue
		}

		img, err := screenshot.CaptureRect(displays[i].bounds)

		if err != nil {
			log.Fatal("Unable to get bounds of requested display", err)
		}

		fileName := CreateDisplayName()

		file, err := os.Create(fileName)

		if err != nil {
			log.Fatal("Unable to create file", err)
		}

		defer file.Close()

		png.Encode(file, img)

	}
}

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("Quick ScreenShot")

	myWindow.Resize(fyne.NewSize(300, 200))

	displays := GetDisplays()

	displaySelector := widget.NewSelect(GetDisplayNames(displays), func(value string) {
		fmt.Println("Screen " + value + " selected")
	})

	screenshotButton := widget.NewButton("Take screenshot", func() {
		takeScreenshot(displays, displaySelector.Selected)
	})

	displayOptionsListContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), displaySelector, layout.NewSpacer())

	takeScreenShotButtonContainer := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), screenshotButton, layout.NewSpacer())

	myWindow.SetContent(container.New(layout.NewVBoxLayout(), layout.NewSpacer(), displayOptionsListContainer, layout.NewSpacer(), takeScreenShotButtonContainer, layout.NewSpacer()))

	myWindow.ShowAndRun()
}
